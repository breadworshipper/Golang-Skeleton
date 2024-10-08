package seeds

import (
	"fmt"
	"math/rand"
	"mm-pddikti-cms/internal/adapter"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
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
		s.usersSeed(total)
	// case "histories":
	// 	s.historiesSeed()
	// case "announcements":
	// 	s.announcementsSeed()
	// case "activities":
	// 	s.activitiesSeed()
	// case "all":
	// 	s.usersSeed(total)
	// 	s.historiesSeed()
	// 	s.announcementsSeed()
	// 	s.activitiesSeed()
	// case "delete-all":
	// 	s.deleteAll()
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

// users
func (s *Seed) usersSeed(total int) {

	tx := s.db.Begin()
	if tx.Error != nil {
		log.Info().Msg("Failed to begin transaction")
		return
	}

	// Constructing user data
	var users []string
	for i := 1; i <= total; i++ {
		fullname := fmt.Sprintf("User%d", i)          // Example name
		username := fmt.Sprintf("user%d", i)          // Example username
		email := fmt.Sprintf("user%d@example.com", i) // Example email
		password := "password"                        // Placeholder password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		role := "admin" // Placeholder role
		users = append(users, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", fullname, username, email, hashedPassword, role))
	}

	// Construct the query
	query := fmt.Sprintf("INSERT INTO users (full_name, username, email, password, role) VALUES %s",
		strings.Join(users, ", "))

	// Execute the INSERT query
	err := tx.Exec(query).Error
	if err != nil {
		log.Info().Msg("users table seed failed: " + err.Error())
		tx.Rollback() // Rollback on error
		return
	}

	// If no error, commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Info().Msg("Commit failed: " + err.Error())
		return
	}
	log.Info().Msg("users table seed success")
}
