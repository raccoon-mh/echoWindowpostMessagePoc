package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestData struct {
	AccessToken   string            `json:"accessToken"`
	WorkspaceInfo map[string]string `json:"workspaceInfo"`
	ProjectInfo   map[string]string `json:"projectInfo"`
	RequestUrl    string            `json:"requestUrl"`
}

type ResponseData struct {
	RequestUrl string `json:"requestUrl"`
}

func main() {
	e := echo.New()

	e.GET("/getinfo", func(c echo.Context) error {
		return c.File("static/getInfo.html")
	})

	e.POST("/gettarget", func(c echo.Context) error {
		var reqData RequestData
		if err := c.Bind(&reqData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// 이곳에서 필요한 인증절차를 진행합니다.

		response := ResponseData{
			RequestUrl: reqData.RequestUrl,
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/target", func(c echo.Context) error {
		return c.File("static/target.html")
	})

	e.Logger.Fatal(e.Start(":8888"))
}
