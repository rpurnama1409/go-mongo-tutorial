package main

import (
	"fmt"
	"go-mongo-tutorial/config"
	"go-mongo-tutorial/src/modules/profile/model"
	"go-mongo-tutorial/src/modules/profile/repository"
	"time"
)

func main() {
	fmt.Println("Go Mongo DB")
	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profiles")

	//saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)
	//getProfile("02", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "02"
	p.FirstName = "ekoroger"
	p.LastName = "Einsten"
	p.Email = "albert@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved")
	}
}

// update profile
func updateProfile(ProfileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "1"
	p.FirstName = "John"
	p.LastName = "Doel Sumbang"
	p.Email = "john@gmail.com"
	p.Password = "123456"
	p.UpdatedAt = time.Now()

	err := ProfileRepository.Update("1", &p)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error")
	} else {
		fmt.Println("Profile updated")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("02")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(profile)
	}
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(profiles)
	}
}
