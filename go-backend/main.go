
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    "fmt"
)

type Member struct {
    PK         uint       `gorm:"primary_key" json:"pk"`
    Username   string     `json:"username"`
    CreateTime time.Time  `json:"create_time"`
}

type BorrowFee struct {
    PK         uint       `gorm:"primary_key" json:"pk"`
    MemberFK   uint       `json:"member_fk"`
    Type       int        `json:"type"`
    BorrowFee  float64    `json:"borrow_fee"`
    CreateTime time.Time  `json:"create_time"`
}

func main() {
    // Replace with your MySQL connection string
    db, err := gorm.Open("mysql", "root:root@tcp(mysql-container:3306)/database_name?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Member{}, &BorrowFee{})

    r := gin.Default()

    // Enable CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    // GET all members
    r.GET("/members", func(c *gin.Context) {
        var members []Member
        db.Find(&members)
        c.JSON(200, members)
    })

    // GET member by ID
    r.GET("/member/:id", func(c *gin.Context) {
        var member Member
        id := c.Param("id")
        if err := db.First(&member, id).Error; err != nil {
            c.JSON(404, gin.H{"error": "Member not found"})
            return
        }

        // Query the total transaction amount for the past year
        var totalFee float64
        if err := db.Model(&BorrowFee{}).
            Where("member_fk = ? AND create_time >= ?", id, time.Now().AddDate(0, 0, -365)).
            Select("sum(borrow_fee)").
            Row().Scan(&totalFee); err != nil {
            // Set total fee to 0 if error occurs
            totalFee = 0
        }

        // Prepare the response with customer information and transaction amount
        response := gin.H{
            "pk":               member.PK,
            "username":         member.Username,
            "create_time":      member.CreateTime,
            "total_fee_last_year": totalFee,
        }

        // Return the response
        c.JSON(200, response)
    })

    // PUT update member by ID
    r.PUT("/member/:id", func(c *gin.Context) {
        var member Member
        id := c.Param("id")
        if err := db.First(&member, id).Error; err != nil {
            c.JSON(404, gin.H{"error": "Member not found"})
            return
        }

        var data map[string]interface{}
        if err := c.BindJSON(&data); err != nil {
            c.JSON(400, gin.H{"error": "Invalid JSON"})
            return
        }

        db.Model(&member).Update("username", data["username"])
        c.JSON(200, gin.H{"message": "Member updated successfully"})
    })

    // POST create new member
    r.POST("/member", func(c *gin.Context) {
        var data map[string]interface{}
        if err := c.BindJSON(&data); err != nil {
            c.JSON(400, gin.H{"error": "Invalid JSON"})
            return
        }

        username, ok := data["username"].(string)
        if !ok || username == "" {
            c.JSON(400, gin.H{"error": "Username is required"})
            return
        }

        member := Member{
            Username: username,
            CreateTime: time.Now(),
        }
        db.Create(&member)
        c.JSON(200, gin.H{"message": "Member created successfully", "pk": member.PK})
    })

    // GET transaction records for a specific customer within a specified time range
    r.GET("/transactions/:memberID", func(c *gin.Context) {
        memberID := c.Param("memberID")

        // Parse startTime and endTime from query string
        startTimeStr := c.Query("startTime")
        endTimeStr := c.Query("endTime")

        // Convert startTime and endTime to time.Time
        startTime, err := time.Parse(time.RFC3339, startTimeStr)
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid startTime"})
            return
        }
        endTime, err := time.Parse(time.RFC3339, endTimeStr)
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid endTime"})
            return
        }

        var transactions []BorrowFee
        if err := db.Where("member_fk = ? AND create_time BETWEEN ? AND ?", memberID, startTime, endTime).Find(&transactions).Error; err != nil {
            c.JSON(500, gin.H{"error": "Failed to retrieve transactions"})
            return
        }

        // Log transactions for debugging
        for i, transaction := range transactions {
            fmt.Printf("Transaction %d: %+v\n", i+1, transaction)
        }
        c.JSON(200, transactions)
    })

    r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
