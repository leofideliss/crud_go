package commands

import (
	"context"
	"crud_go/internal/domain"
	"fmt"
	"strconv"
	"sync"
	"time"
    //    "log"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/go-redsync/redsync/v4"
	"github.com/go-redis/redis/v8"                        
	redigo "github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var lock sync.Mutex
var mutex *redsync.Mutex
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualizada o limite de uma categoria",
	Run: func(cmd *cobra.Command, args []string) {
        start := time.Now()
        repeat, _ := cmd.Flags().GetString("repeat")
        loop := 1
        if repeat != "" {
            repeatInt , err := strconv.Atoi(repeat)
            if err != nil{
                fmt.Println("O valor não pode ser convertido")
            }
            loop = repeatInt
        }
        
        done := make(chan bool)
        for i := 0 ; i < loop ; i++{
        	go func() {
				updateCategories(cmd)
				done <- true 
			}()
        }

        for i := 0; i < loop; i++ {
			<-done
		}

        elapsed := time.Since(start);
        fmt.Printf("Tempo de execução: %s\n" , elapsed);
	},
}

func init() {
	updateCmd.Flags().StringP("id", "d", "", "Id da categoria que será atualizada")
	updateCmd.Flags().StringP("qty", "q", "", "Valor passado para atualizar a categoria")
	updateCmd.Flags().StringP("repeat", "r", "", "Quantidade de interações")


    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    pool := redigo.NewPool(client)
    rs := redsync.New(pool)
    mutex = rs.NewMutex("db-lock")
}

func updateCategories(cmd *cobra.Command){
    fmt.Println("Atualizando valores...")
    id, _ := cmd.Flags().GetString("id")
    qty, _ := cmd.Flags().GetString("qty")

    if id == ""{
        fmt.Println("Informar o Id da categoria")
    }

    objId , _ := primitive.ObjectIDFromHex(id)
    lock.Lock()
    defer lock.Unlock()


    // Tenta adquirir o lock
    // if err := mutex.Lock(); err != nil {
    //     log.Fatal(err)
    //}


    result, _ := CategoryService.ReadCategories(context.Background(),objId)
    
    if qty == ""{
        fmt.Println("Informe um valor")
    }
    
    qtyInt , _ := strconv.Atoi(qty)
    
    switch qtyInt > 0 {
    case true :
        result.Limit += qtyInt
    case false :
        result.Limit -= qtyInt
    }
    
    category :=  domain.Category{Name:result.Name , Tag:result.Tag , Limit:result.Limit,}
    _, errUpdate := CategoryService.UpdateCategory(context.Background(), &category ,objId)
    if errUpdate != nil {
        fmt.Println("Erro ao atualizar")
    }

    
	// Libera o lock
	// if ok, err := mutex.Unlock(); !ok || err != nil {
	// 	log.Fatalf("Erro ao liberar lock: %v", err)
    //	}
    
}
