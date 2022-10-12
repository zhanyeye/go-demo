package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	r := gin.Default()
	r.GET("/api/fib/:num/*action", func(c *gin.Context) {
		num, err := strconv.Atoi(c.Param("num"))
		if err != nil || num > 20 || num < 0 {
			c.String(http.StatusBadRequest, c.Param("num")+" 输入校验不通过")
			panic(err)
		}
		c.String(http.StatusOK, "第 "+c.Param("num")+" 个斐波那契数是："+strconv.Itoa(fib(num)))

	})
	//默认为监听8080端口
	r.Run(":8000")
}
