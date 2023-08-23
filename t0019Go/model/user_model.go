package model

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sharin-sushi/0019testGin.git/crypto"

	_ "github.com/go-sql-driver/mysql"
)

type Test01_user struct { //dbに対してはtable名 小文字かつ複数形に自動変換
	// gorm.Model
	UserId   int // `gorm:"primaryKey"` //→dbに対してはスネークケースに自動変換
	Password string
}

func (u *Test01_user) LoggedIn() bool {
	// return u.ID != 0 元々これ
	return u.UserId != 0
}

func Signup(user_id, password string) (*Test01_user, error) {
	var test01_user Test01_user
	// test01_user := Test01_user だとダメ　gorm公式はこの書き方
	// test01_user := Test01_user{}最初はこうだった

	fmt.Printf("Sinup関数の受け取った値:user_id=%v, password=%v \n", user_id, password)
	// fmt.Printf("test01_user=%v \n", test01_user)

	// test01_user.UserId = 2
	// test01_user.Password = "2"
	// fmt.Printf("test01_user=%v, user_id=%s \n", test01_user, user_id)

	Db.Where("user_id = ?", user_id).First(&test01_user.UserId) //.Firts 1件だけ()内に格納→DBからEnmpty setが返ってきてるはず

	fmt.Printf("userId=%v, password=%v \n", test01_user.UserId, test01_user.Password) //{0, }期待→ok

	if test01_user.UserId != 0 {
		err := errors.New("登録済みのUserIdです。")
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("test01_user=%vです。 \n", test01_user)

	encryptPw, err := crypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化処理でエラー発生：", err)
		return nil, err
	}
	fmt.Printf("encryptPw=%v \n", encryptPw)

	user_idINT, _ := strconv.Atoi(user_id)
	test01_user = Test01_user{UserId: user_idINT, Password: encryptPw}
	fmt.Printf("test01_user=%v \n", test01_user)
	Db.Create(&test01_user)
	fmt.Printf("a \n")
	return &test01_user, nil
}

func Login(userId, password string) (*Test01_user, error) {
	user := Test01_user{}
	Db.Where("user_id = ?", userId).First(&user)
	if user.UserId == 0 {
		err := errors.New("UserIdが一致するユーザーが存在しません。")
		fmt.Println(err)
		return nil, err
	}

	err := crypto.CompareHashAndPassword(user.Password, password)
	if err != nil {
		fmt.Println("パスワードが一致しませんでした。：", err)
		return nil, err
	}

	return &user, nil
}
