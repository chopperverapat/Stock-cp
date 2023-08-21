# Middleware และ Interceptor

Middleware เป็นตัวกลางที่ใช้ในการประมวลผล HTTP requests ก่อนที่จะถึง handler หลักของเรา (controller) หรือหลังจากที่ handler ทำงานเสร็จแล้วก่อนที่จะส่ง response กลับไปให้ client ในการใช้งานจริง Middleware สามารถใช้งานเพื่อปรับแต่งเช่น เช็คการรับส่งข้อมูล ตรวจสอบการรับส่ง Token, ทำการประมวลผลก่อนเข้าสู่การจัดการหลัก หรือทำการล็อกอิน ฯลฯ

Interceptor นั้นเป็นส่วนหนึ่งของ Middleware ที่ทำหน้าที่ตรวจสอบหรือประมวลผลบางอย่างก่อนที่ HTTP request จะถึง handler หลักของเรา โดย Interceptor จะสามารถควบคุมการทำงานของเราก่อนถึงจะถึงมือในการจัดการของ Controller หรือ Handler หลัก

## ตัวอย่าง Middleware แบบง่าย - ตรวจสอบ Token และสิทธิ์การเข้าถึง

```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // ตรวจสอบ Token หรือการรับส่งข้อมูลตามที่ต้องการ
        
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort() // หยุดการทำงานของ Middleware นี้
            return
        }
        
        // ตรวจสอบสิทธิ์การเข้าถึงตามที่ต้องการ
        // ตัวอย่างเช่น ตรวจสอบว่าผู้ใช้มีสิทธิ์เข้าถึงหน้าที่กำลังจะเข้าถึงหรือไม่
        // ถ้าไม่มีสิทธิ์ให้ทำการหยุดการทำงานของ Middleware และส่งข้อความ Unauthorized
        // เป็นต้น
        // ...
        
        // ถ้าผ่านการตรวจสอบทั้งหมด ให้ผ่านการทำงานไปยัง Handler หลัก
        c.Next()
    }
}

func SetupAuthenApi(router *gin.Engine) {
    authenAPI := router.Group("/api/v2")
    authenAPI.Use(AuthMiddleware()) // ใช้ Middleware ในกลุ่มนี้
    {
        authenAPI.POST("/login", login)
        authenAPI.POST("/register", register)
    }
}

```
## ตัวอย่างที่ 2

```go

// create interepter
func AuthenMiddleware(c *gin.Contest){
    // query from clien sent request
    // http://brabra/api/v2/?token=1234 => string query
    token := c.Query("token")
    if token = "1234" {
        //ให้ทำต่อ ไป function hendle ได้
        c.Next()
    }else {
        c.JSON(http.StatusUnthorizeds,gin.H{"error": "invalid token"})
        // ให้หยุดการทำงาน
        c.Abort()
    }
}

func SetupProductApi(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
        // Interceptor มาขั้นก่อนจะถึง function handle ซึ่ง Interceptor เป็นส่วนึงของ middleware
		productAPI.GET("/product", AuthenMiddleware,getproduct)
		productAPI.POST("/product", createproduct)
	}
}

```
