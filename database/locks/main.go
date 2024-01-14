package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Bloqueio Otimista (Optimistic Locking):

Definição: Uma técnica de controle de concorrência usada em ambientes onde conflitos são considerados raros, mas possíveis.
Não bloqueia o recurso quando lido, mas verifica se houve alterações antes de atualizar.

Mecanismo: Registra uma versão ou timestamp quando um registro é lido.
Quando uma atualização é tentada, o sistema verifica se a versão ou timestamp mudou desde a leitura original.
Se mudou, isso indica que outro processo atualizou o registro, e a atualização atual é rejeitada ou requer intervenção.

Uso Ideal: Ambientes de alta concorrência onde as atualizações são relativamente raras e o custo de rejeitar e retrabalhar uma transação é baixo.
Exemplos incluem muitos sistemas web e aplicações onde a leitura é muito mais frequente que a escrita.
*/

/*
Bloqueio Pessimista (Pessimistic Locking):

Definição: Uma técnica de controle de concorrência que bloqueia um recurso quando ele é lido para atualização,
impedindo outros processos de acessar o mesmo recurso para escrita ou atualizações.

Mecanismo: Quando um registro é lido para atualização, ele é bloqueado e outros processos não podem modificar esse registro até que o bloqueio seja liberado.
Isso previne conflitos ao custo de reduzir a concorrência.

Uso Ideal: Situações onde é esperado que as transações mantenham bloqueios por um período mais longo,
ou em sistemas onde os conflitos de atualização são frequentes e o custo de lidar com transações rejeitadas é alto.
Exemplos incluem processamento de transações financeiras ou operações que envolvem múltiplos passos interdependentes.
*/

/*
O bloqueio otimista favorece a performance em cenários de alta leitura e baixa escrita,
enquanto o bloqueio pessimista é mais robusto em situações onde as escritas são mais críticas e conflituosas.
*/

type Item struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:itens_categories;"` // many2many relationship
	gorm.Model
}

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Items []Item `gorm:"many2many:itens_categories;"` // many2many relationship
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&Item{}, &Category{})

	// Pessimistic Locking using FOR UPDATE clause with transaction
	tx := db.Begin()
	var c Category
	err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	c.Name = "Home"
	tx.Debug().Save(&c)
	tx.Commit()

}
