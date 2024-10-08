package seeds

import (
	"fmt"
	"math/rand"
	"mm-pddikti-cms/internal/adapter"
	user_entity "mm-pddikti-cms/internal/module/user/entity"
	"mm-pddikti-cms/pkg"
	"os"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Seed struct.
type Seed struct {
	db *gorm.DB
}

// var uuids = []string{
// 	"550e8400-e29b-41d4-a716-446655440000",
// 	"550e8400-e29b-41d4-a716-446655440001",
// 	"550e8400-e29b-41d4-a716-446655440002",
// 	"550e8400-e29b-41d4-a716-446655440003",
// 	"550e8400-e29b-41d4-a716-446655440004",
// 	"550e8400-e29b-41d4-a716-446655440005",
// 	"550e8400-e29b-41d4-a716-446655440006",
// 	"550e8400-e29b-41d4-a716-446655440007",
// 	"550e8400-e29b-41d4-a716-446655440008",
// 	"550e8400-e29b-41d4-a716-446655440009",
// 	"550e8400-e29b-41d4-a716-446655440010",
// 	"550e8400-e29b-41d4-a716-446655440011",
// 	"550e8400-e29b-41d4-a716-446655440012",
// 	"550e8400-e29b-41d4-a716-446655440013",
// }

var uuidBase = "550e8400-e29b-41d4-a716-4466554400"

// NewSeed return a Seed with a pool of connection to a dabase.
func newSeed(db *gorm.DB) Seed {
	return Seed{
		db: db,
	}
}

func Execute(db *gorm.DB, table string, total int) {
	seed := newSeed(db)
	seed.run(table, total)
}

// Run seeds.
func (s *Seed) run(table string, total int) {
	switch table {
	case "users":
		s.superAdminUserSeed()
		s.adminUsersSeed(total)
	case "histories":
		s.historiesSeed()
	case "announcements":
		s.announcementsSeed()
	case "activities":
		s.activitiesSeed()
	case "all":
		s.superAdminUserSeed()
		s.adminUsersSeed(total)
		s.historiesSeed()
		s.announcementsSeed()
		s.activitiesSeed()
	case "delete-all":
		s.deleteAll()
	default:
		log.Warn().Msg("No seed to run")
	}

	if table != "" {
		log.Info().Msg("Seed ran successfully")
		log.Info().Msg("Exiting ...")
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
		os.Exit(0)
	}
}

func (s *Seed) deleteAll() {

}

func (s *Seed) historiesSeed() {
	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}

	var histories []string
	for i := 1; i <= 15; i++ {
		user_id := fmt.Sprintf("%s%d", uuidBase, i)
		actions := []string{"create", "update", "delete"}
		action := actions[rand.Intn(len(actions))]
		entities := []string{"announcement", "history", "activity"}
		entity_name := entities[rand.Intn(len(entities))]
		entity_id := fmt.Sprintf("%s%d", uuidBase, i)

		histories = append(histories, fmt.Sprintf("(%v, %v, %v, %v)", user_id, action, entity_name, entity_id))
	}

	query := fmt.Sprintf("INSERT INTO histories (user_id, action, entity_name) VALUES %s",
		strings.Join(histories, ", "))

	err := tx.Exec(query).Error
	if err != nil {
		log.Info().Msg("histories table seed failed: " + err.Error())
		tx.Rollback() // Rollback on error
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}

	log.Info().Msg("histories table seed success")
}

func (s *Seed) activitiesSeed() {
	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}

	var activities []string
	for i := 1; i <= 15; i++ {
		title := fmt.Sprintf("Activity Title %d", i)
		slug := fmt.Sprintf("activity-title-%d", i)
		content := fmt.Sprintf("this is content for activity-%d", i)
		thumbnail := fmt.Sprintf("This is description for title %d", i)

		activities = append(activities, fmt.Sprintf("(%v, %v, %v, %v)", title, slug, content, thumbnail))
	}

	query := fmt.Sprintf("INSERT INTO activities (title, slug, content, thumbnail) VALUES %s",
		strings.Join(activities, ","))

	err := tx.Exec(query).Error
	if err != nil {
		log.Info().Msg("activities table seed failed: " + err.Error())
		tx.Rollback() // Rollback on error
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}

	log.Info().Msg("activities table seed success")
}

func (s *Seed) announcementsSeed() {
	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}
	var announcements []string
	for i := 1; i <= 15; i++ {
		title := fmt.Sprintf("Announcement Title %d", i)
		slug := fmt.Sprintf("announcement-title-%d", i)
		link := fmt.Sprintf("http://link-to-title-%d", i)
		description := fmt.Sprintf("This is description for announcement %d", i)

		announcements = append(announcements, fmt.Sprintf("(%v, %v, %v, %v)", title, slug, link, description))
	}

	query := fmt.Sprintf("INSERT INTO announcements (title, slug, link, description) VALUES %s",
		strings.Join(announcements, ","))

	err := tx.Exec(query).Error
	if err != nil {
		log.Info().Msg("announcement table seed failed: " + err.Error())
		tx.Rollback() // Rollback on error
		return
	}
	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}

	log.Info().Msg("announcement table seed success")
}

func (s *Seed) superAdminUserSeed() {
	var count int64
	s.db.Model(&user_entity.User{}).Where("role = ?", user_entity.RoleSuperAdmin).Count(&count)
	if count > 0 {
		log.Info().Msg("Super admin already seeded")
		return
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Info().Msg("Failed to generate UUID")
		return
	}

	password, err := pkg.HashPassword("superadmin")
	if err != nil {
		log.Info().Msg("Failed to hash password")
		return
	}

	err = tx.Create(&user_entity.User{
		ID:       newUUID,
		FullName: "Super Admin",
		Username: "superadmin",
		Email:    "superadmin@cms.pddikti.kemdikbud.go.id",
		Password: password,
		Role:     user_entity.RoleSuperAdmin,
	}).Error
	if err != nil {
		log.Info().Msg("users table (super-admin) seed failed: " + err.Error())
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}

	log.Info().Msg("users table (super-admin) seed success")
}

func (s *Seed) adminUsersSeed(total int) {
	type UserFaker struct {
		FullName string `faker:"name"`
		Username string `faker:"username,unique"`
		Email    string `faker:"email,unique"`
		Password string `faker:"password"`
	}

	var users []user_entity.User

	for i := 1; i <= total; i++ {
		fakeUser := UserFaker{}
		err := faker.FakeData(&fakeUser)
		if err != nil {
			fmt.Println(err)
			continue
		}

		newUUID, err := uuid.NewUUID()
		if err != nil {
			log.Info().Msg("Failed to generate UUID")
			return
		}

		user := user_entity.User{
			ID:        newUUID,
			FullName:  fakeUser.FullName,
			Username:  fakeUser.Username,
			Email:     fakeUser.Email,
			Password:  fakeUser.Password,
			Role:      user_entity.RoleAdmin,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		users = append(users, user)
	}
	faker.ResetUnique()

	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}

	err := tx.Create(&users).Error
	if err != nil {
		log.Info().Msg("users table (admin) seed failed: " + err.Error())
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}

	log.Info().Msg("users table (admin) seed success")
}
