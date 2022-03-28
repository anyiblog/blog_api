package api

//根接口

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/service"
	"anyiblog/util"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

// Ping 服务器通达性
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "ok",
	})
}

// DeleteFile 根据图片链接删除图片
func DeleteFile(c *gin.Context) {
	type delImg struct {
		ImgUrl string `json:"imgUrl"`
	}
	delImgPrams := &delImg{}
	if err := c.ShouldBindJSON(&delImgPrams); err == nil {
		filePath := delImgPrams.ImgUrl                                      //filePath带域名，删除数据库需要
		fileName := strings.Replace(filePath, os.Getenv("COS_Url"), "", -1) //fileName不带，删除Cos需要
		cosErr := service.DeleteCosFile(fileName)                           //删除Cos文件
		if cosErr == nil {
			if model.DeleteImg(filePath) {
				c.JSON(200, serializer.Response{
					Code: serializer.SystemOk,
					Msg:  "删除文件成功",
				})
			} else {
				c.JSON(200, serializer.Response{
					Code: serializer.SystemError,
					Msg:  "删除文件失败",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Code: serializer.SystemError,
				Msg:  "删除失败：" + cosErr.Error(),
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: serializer.SystemError,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

type Img struct {
	ImgId  string `json:"img_id"`
	ImgUrl string `json:"img_url"`
}
type ResImgInfo struct {
	ImgId   string    `json:"img_id"`   // UUID
	ImgName string    `json:"img_name"` // UUID + 文件后缀名
	ImgData io.Reader `json:"img_data"` // 二进制数据
}

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	fmt.Println(c.ContentType())
	if c.ContentType() == "application/json" { //json 批量上传
		type UploadFileList struct {
			List []string `json:"list" binding:"required"`
		}
		var uploadFileList UploadFileList
		if err := c.ShouldBindJSON(&uploadFileList); err == nil {
			paramUrlList := uploadFileList.List //图片链接列表
			var resFileList []Img               //上传成功后CosUrl文件列表
			for i := range paramUrlList {
				imgItemUrl := paramUrlList[i]
				imgItemInfo, errInfo := getImg(imgItemUrl)
				if errInfo != nil {
					c.JSON(200, serializer.Response{
						Code: serializer.SystemError,
						Msg:  errInfo.Error(),
					})
					return
				}
				cosRes := uploadFile(imgItemInfo.ImgData, imgItemInfo.ImgName)
				if cosRes.Code == serializer.SystemOk { // Cos文件上传成功
					resFileList = append(resFileList, Img{
						ImgId:  imgItemInfo.ImgId,
						ImgUrl: cosRes.Data.(string),
					})
				} else {
					c.JSON(200, cosRes)
					return
				}
			}
			for i := range resFileList { // 将图片记录插入数据库
				model.CreateImgByUrl(resFileList[i].ImgUrl, resFileList[i].ImgId)
			}
			c.JSON(200, serializer.Response{
				Code: 0,
				Msg:  "上传成功",
				Data: resFileList,
			})
		} else {
			c.JSON(200, serializer.Response{
				Code: 1,
				Msg:  "参数错误",
			})
		}
	} else if c.ContentType() == "multipart/form-data" { //form-data 单文件上传
		form, _ := c.MultipartForm()
		// 获取所有图片
		files := form.File["file"]
		var resFileList []Img         //上传成功后CosUrl文件列表
		var uploadOkFileList []string //已上传成功的文件名
		for _, file := range files {
			fData, _ := file.Open() //io.reader
			// 文件名为：UUID + 后缀
			fileId := util.GenerateUUID()
			fileExt := path.Ext(file.Filename)
			fileName := fileId + fileExt

			res := uploadFile(fData, fileName)
			if res.Code == serializer.SystemOk { // 当前文件上次成功
				resFileList = append(resFileList, Img{
					ImgId:  fileId,
					ImgUrl: res.Data.(string),
				})
				uploadOkFileList = append(uploadOkFileList, fileName)
			} else {
				c.JSON(200, res)
				for i := range uploadOkFileList {
					_ = service.DeleteCosFile(uploadOkFileList[i])
				}
			}
			_ = fData.Close()
		}
		for i := range resFileList { //将图片记录插入数据库
			model.CreateImgByUrl(resFileList[i].ImgUrl, resFileList[i].ImgId)
		}
		c.JSON(200, serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "文件上传成功",
			Data: resFileList,
		})
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求类型错误",
		})
	}
}

//文件上传具体实现（二进制流，文件名）
func uploadFile(fData io.Reader, fileName string) serializer.Response {
	fileUrl, fileErr := service.CosUpload(fData, fileName)
	if fileErr != nil {
		res := serializer.Response{
			Code: serializer.SystemError,
			Msg:  fileErr.Error(),
		}
		return res
	} else {
		res := serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "文件上传成功",
			Data: fileUrl,
		}
		return res
	}
}

//保存文件
func getImg(url string) (ResImgInfo, error) {
	//获取远端图片
	resp, err := http.Get(url)
	if err != nil {
		return ResImgInfo{}, errors.New("连接图片服务器错误")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	imgType := resp.Header.Get("Content-Type")
	if imgType == "image/png" || imgType == "image/jpeg" || imgType == "image/gif" {
		fileExt := strings.Split(imgType, "/")
		// 读取获取的[]byte数据
		imgData, _ := ioutil.ReadAll(resp.Body)

		fileId := util.GenerateUUID()
		fileName := fileId + "." + fileExt[1]
		resImgInfo := ResImgInfo{
			ImgId:   fileId,
			ImgName: fileName,
			ImgData: bytes.NewReader(imgData),
		}
		return resImgInfo, nil
	} else {
		// URL不是正确的图片链接
		return ResImgInfo{}, errors.New("URL不是正确的图片链接")
	}
}
