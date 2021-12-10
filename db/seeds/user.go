package seeds

import (
	"tubespbbo/db"
	"tubespbbo/hashing"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedUser() {
	var user []model.User
	var count int64
	db.Orm.Model(&user).Count(&count)

	if count > 0 {
		return
	}

	users := make([]model.User, 7)
	users[0] = model.User{
		Username: "Jojoan",
		Password: hashing.HashAndSalt([]byte("joan123")),
		Name:     "Eirenika Joanna Grace",
		Dob:      "2000-10-02",
		Role:     0,
		Address:  "Jln. Dimana-mana hatiku senang no.1",
		Phone:    "081122334455",
		Email:    "joanjo@email.com",
	}
	users[1] = model.User{
		Username: "JeniJenJen",
		Password: hashing.HashAndSalt([]byte("jeni123")),
		Name:     "Jenifer Florentina",
		Dob:      "2001-06-18",
		Role:     1,
		Address:  "Jln. Dimana-mana hatiku gembira no.1",
		Phone:    "081166887799",
		Email:    "jeniakatina@email.com",
	}
	users[2] = model.User{
		Username: "Marimar",
		Password: hashing.HashAndSalt([]byte("maria123")),
		Name:     "Maria Vabiolla",
		Dob:      "2001-10-12",
		Role:     2,
		Address:  "Jln. Dimana-mana hatiku bahagia no.1",
		Phone:    "081100337788",
		Email:    "mariaAjah@email.com",
	}
	users[3] = model.User{
		Username: "Enjeljel",
		Password: hashing.HashAndSalt([]byte("enjel123")),
		Name:     "Tridia Enjeliani",
		Dob:      "2001-10-04",
		Role:     3,
		Address:  "Jln. Dimana-mana hatiku tertawa no.1",
		Phone:    "081122998866",
		Email:    "jeniakatina@email.com",
	}
	users[4] = model.User{
		Username: "Melpinz",
		Password: hashing.HashAndSalt([]byte("melvin123")),
		Name:     "Melvin Sebastian",
		Dob:      "2001-07-21",
		Role:     4,
		Address:  "Jln. Dimana-mana hatiku hanya dia no.1",
		Phone:    "0811229988",
		Email:    "melpinzs@email.com",
	}
	users[5] = model.User{
		Username: "Boedi123",
		Password: hashing.HashAndSalt([]byte("budi123")),
		Name:     "Budiarto",
		Dob:      "2000-12-10",
		Role:     0,
		Address:  "Jln. Sisa kenangan masa lalu no.1",
		Phone:    "089822234411",
		Email:    "boedi123@email.com",
	}
	users[6] = model.User{
		Username: "Nichoco",
		Password: hashing.HashAndSalt([]byte("ndutdut123")),
		Name:     "Nicholas Ndut",
		Dob:      "1999-12-12",
		Role:     0,
		Address:  "Jln. Dihati enjel no.1",
		Phone:    "081234567899",
		Email:    "ndudut@email.com",
	}

	_ = db.Orm.Create(&users)
}
