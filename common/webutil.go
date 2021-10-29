package webutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nbit99/open_base/model"
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/blocktree/openwallet/v2/owtp"
	"github.com/godaddy-x/jorm/consul"
	"github.com/godaddy-x/jorm/exception"
	"github.com/godaddy-x/jorm/log"
	"github.com/godaddy-x/jorm/util"
	"io/ioutil"
	"net/http"
	"reflect"
	"unsafe"
)

// map参数转成req对象
func ToReqDto(ctx *owtp.Context, args interface{}) error {
	param := ctx.Params().Map()
	if param == nil || len(param) == 0 {
		return util.Error("参数异常")
	}
	var fields []map[string]reflect.Type
	var values []string
	tof := reflect.TypeOf(args).Elem()
	for i := 0; i < tof.NumField(); i++ {
		field := tof.Field(i)
		k, err := util.GetJsonTag(field)
		if k == "" || err != nil {
			continue
		}
		v, ok := param[k]
		if !ok {
			continue
		}
		fields = append(fields, map[string]reflect.Type{k: field.Type})
		values = append(values, v.String())
	}
	page := make([]int64, 0)
	pageNo, ok := param["pageNo"]
	if ok {
		i64, err := util.StrToInt64(pageNo.String())
		if err != nil {
			return err
		}
		page = append(page, i64)
	}
	pageSize, ok := param["pageSize"]
	if ok {
		i64, err := util.StrToInt64(pageSize.String())
		if err != nil {
			return err
		}
		page = append(page, i64)
	}
	if str, err := dataToMap(fields, values, page, ctx.GetSession("userId")); err != nil {
		log.Error("请求参数dataToMap转换失败", 0, log.Any("fields", fields), log.Any("values", values), log.AddError(err))
		return err
	} else if err := util.JsonToObject(str, args); err != nil {
		log.Error("请求参数JsonToObject转换失败", 0, log.String("args", str), log.AddError(err))
		return err
	}
	return nil
}

func GetRpc() (*consul.ConsulManager, error) {
	consulx, err := new(consul.ConsulManager).Client()
	if err != nil {
		return nil, err
	}
	return consulx, nil
}

//responseSuccess 成功结果响应
func ResponseSuccess(ctx *owtp.Context, result interface{}) {
	ctx.Response(result, owtp.StatusSuccess, "success")
}

//responseError 失败结果响应
func ResponseError(ctx *owtp.Context, err error) {
	ctx.Response(nil, owtp.ErrCustomError, err.Error())
}

//responseError 失败结果响应
func ResponseOwError(ctx *owtp.Context, err error) {
	ex := Catch(err)
	if ex.Code == 0 {
		ctx.Response(nil, owtp.ErrCustomError, err.Error())
	} else {
		ctx.Response(nil, ex.Code, util.AddStr(ex.ErrMsg, " >>> ", ex.ExtMsg))
	}
}

//responseError 失败结果响应
func ResponseNewError(ctx *owtp.Context, err error) {
	e := ex.Catch(err)
	ctx.Response(nil, uint64(e.Code), e.Msg)
}

// 结果集根据对象字段类型填充到map实例
func dataToMap(fieldArray []map[string]reflect.Type, values []string, page []int64, userId ...interface{}) (string, error) {
	result := make(map[string]interface{})
	for i := 0; i < len(fieldArray); i++ {
		for f, t := range fieldArray[i] {
			kind := t.String()
			if kind == "int" {
				i0, err := util.StrToInt(values[i])
				if err != nil {
					return "", util.Error("对象字段[", f, "]转换int失败")
				}
				result[f] = i0
			} else if kind == "int8" {
				i8, err := util.StrToInt8(values[i])
				if err != nil {
					return "", util.Error("对象字段[", f, "]转换int8失败")
				}
				result[f] = i8
			} else if kind == "int16" {
				i16, err := util.StrToInt16(values[i])
				if err != nil {
					return "", util.Error("对象字段[", f, "]转换int16失败")
				}
				result[f] = i16
			} else if kind == "int32" {
				i32, err := util.StrToInt32(values[i])
				if err != nil {
					return "", util.Error("对象字段[", f, "]转换int32失败")
				}
				result[f] = i32
			} else if kind == "int64" {
				i64, err := util.StrToInt64(values[i])
				if err != nil {
					return "", util.Error("对象字段[", f, "]转换int64失败")
				}
				result[f] = i64
			} else if kind == "string" {
				result[f] = values[i]
			} else if kind == "[]string" {
				array := make([]string, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]int" {
				array := make([]int, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]int8" {
				array := make([]int8, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]int16" {
				array := make([]int16, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]int32" {
				array := make([]int32, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]int64" {
				array := make([]int64, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]interface {}" {
				array := make([]interface{}, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "[]string" {
				array := make([]string, 0)
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &array); err != nil {
						return "", err
					}
				}
				result[f] = array
			} else if kind == "map[string]interface {}" {
				mmp := make(map[string]interface{})
				if len(values[i]) > 0 {
					if err := util.JsonToObject(values[i], &mmp); err != nil {
						return "", err
					}
				}
				result[f] = mmp
			} else if kind == "model.CallbackNode" {
				if len(values[i]) > 0 {
					tx := model.CallbackNode{}
					if err := util.JsonToObject(values[i], &tx); err != nil {
						return "", err
					}
					result[f] = tx
				}
			} else if kind == "*openwallet.FeesSupportAccount" {
				if len(values[i]) > 0 {
					tx := &openwallet.FeesSupportAccount{}
					if err := util.JsonToObject(values[i], &tx); err != nil {
						return "", err
					}
					result[f] = tx
				}
			} else {
				return "", util.Error("对象字段[", f, "]转换失败([", kind, "]类型不支持)")
			}
		}
	}
	if len(page) == 2 {
		result["pageNo"] = page[0]
		result["pageSize"] = page[1]
	}
	if userId != nil && userId[0] != nil {
		result["userId"] = userId[0]
	}
	return util.ObjectToJson(result)
}

func ToPost(url string, data interface{}) {
	bytesData, err := json.Marshal(&data)
	if err != nil {
		log.Error("json转换参数异常", 0, log.Any("content", data))
		return
	}
	fmt.Println("请求示例: ")
	fmt.Println(string(bytesData))
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Error("post请求失败", 0, log.String("url", url))
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("a", "123456")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println("响应示例: ")
	fmt.Println(*str)
}

func ToPostTostr(url string, data interface{}) string {
	bytesData, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return *str
}
