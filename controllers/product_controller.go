package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"crud-api/models"
	"strconv"
	"fmt"
    "strings"
)




func GetHistory(c *fiber.Ctx) error {
    db := c.Locals("db").(*gorm.DB)
    var history []models.UserRequest
	db.Order("created_at desc").Limit(20).Find(&history)
	return c.JSON(history)
}



func GetAllEndpoints(c *fiber.Ctx) error {
    endpointList := []string{
        "/",
        "/history",
        "/5/plus/3",
        "/3/minus/5",
        "/12/div/2",
        "/3/into/5",
        "/3/minus/5/plus/8",
        
    }

    additionalText := "These are the available GET endpoints in the application."

    response := struct {
        Endpoints      []string `json:"endpoints"`
        AdditionalText string   `json:"additionalText"`
    }{
        Endpoints:      endpointList,
        AdditionalText: additionalText,
    }

    return c.JSON(response)
}


func Addition(c *fiber.Ctx) error {
    num1, err1 := strconv.Atoi(c.Params("num1"))
    num2, err2 := strconv.Atoi(c.Params("num2"))

    if err1 != nil || err2 != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid numbers"})
    }

    result := num1 + num2
    question := fmt.Sprintf("%d plus %d", num1, num2)

    db := c.Locals("db").(*gorm.DB)
    
    
    userRequest := models.UserRequest{
        Command: "Addition",
        Result:  result,
        Question: question, 
    }

    db.Create(&userRequest)

    return c.JSON(fiber.Map{"result": result})
}


func Subtraction(c *fiber.Ctx) error {
    num1, err1 := strconv.Atoi(c.Params("num1"))
    num2, err2 := strconv.Atoi(c.Params("num2"))

    if err1 != nil || err2 != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid numbers"})
    }

    result := num1 - num2
	question := fmt.Sprintf("%d minus %d", num1, num2)

    db := c.Locals("db").(*gorm.DB)
    
    
    userRequest := models.UserRequest{
        Command: "Subtraction",
        Result:  result,
        Question: question, 
    }

    db.Create(&userRequest)

    return c.JSON(fiber.Map{"result": result})
}

func Multiplication(c *fiber.Ctx) error {
    num1, err1 := strconv.Atoi(c.Params("num1"))
    num2, err2 := strconv.Atoi(c.Params("num2"))

    if err1 != nil || err2 != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid numbers"})
    }

    result := num1 * num2

    
    question := fmt.Sprintf("%d multiplied by %d", num1, num2)

    db := c.Locals("db").(*gorm.DB)
    
    
    userRequest := models.UserRequest{
        Command: "Multiplication",
        Result:  result,
        Question: question, 
    }

    db.Create(&userRequest)

    return c.JSON(fiber.Map{"result": result})
}

func Division(c *fiber.Ctx) error {
    num1, err1 := strconv.Atoi(c.Params("num1"))
    num2, err2 := strconv.Atoi(c.Params("num2"))

    if err1 != nil || err2 != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid numbers"})
    }

    if num2 == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Division by zero is not allowed"})
    }

    result := float64(num1) / float64(num2)

    
    question := fmt.Sprintf("%d divided by %d", num1, num2)

    db := c.Locals("db").(*gorm.DB)
    
    
    userRequest := models.UserRequest{
        Command: "Division",
        Result:  int(result),
        Question: question, 
    }

    db.Create(&userRequest)

    return c.JSON(fiber.Map{"result": result})
}

func ComplexOperation(c *fiber.Ctx) error {
    
    path := c.Params("*")
    segments := strings.Split(path, "/")

    fmt.Printf("Segments: %#v\n", segments) 

    if len(segments) < 3 || len(segments)%2 != 1 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid operation"})
    }


    firstSegment := segments[0]
    result, err := strconv.Atoi(firstSegment)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid operation"})
    }

    
    question := firstSegment

   
    for i := 1; i < len(segments); i += 2 {
        operator := segments[i]
        operandStr := segments[i+1]
        operand, err := strconv.Atoi(operandStr)

        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid operation"})
        }

        switch operator {
        case "plus":
            result += operand
            question += " + " + operandStr
        case "minus":
            result -= operand
            question += " - " + operandStr
        
        default:
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid operator"})
        }
    }

    fmt.Printf("Result: %d, Question: %s\n", result, question) 

    db := c.Locals("db").(*gorm.DB)

    userRequest := models.UserRequest{
        Command:  "Complex Operation",
        Result:   result,
        Question: question,
    }

    db.Create(&userRequest)

    return c.JSON(fiber.Map{"result": result})
}
