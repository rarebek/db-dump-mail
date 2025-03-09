package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Load environment variables from .env file
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// Function to create a PostgreSQL dump
func dumpDatabase() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dumpFile := os.Getenv("DUMP_FILE")

	cmd := exec.Command("pg_dump",
		"-h", host,
		"-p", port,
		"-U", user,
		"--format=plain",
		"--file="+dumpFile,
		dbname)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pg_dump failed: %s, output: %s", err, output)
	}
	return nil
}

// Function to send email with the dump file
func sendEmail() error {
	senderEmail := os.Getenv("MAILER_USERNAME")
	password := os.Getenv("MAILER_PASSWORD")
	receiverEmails := os.Getenv("MAILER_RECEIVERS")
	dumpFile := os.Getenv("DUMP_FILE")

	// Convert comma-separated emails into a slice
	receivers := strings.Split(receiverEmails, ",")

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", receivers...) // Spread to handle multiple recipients
	m.SetHeader("Subject", "PostgreSQL Backup")
	m.SetBody("text/plain", "Find the attached PostgreSQL backup.")
	m.Attach(dumpFile)

	d := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	fmt.Println("Starting PostgreSQL backup...")

	if err := dumpDatabase(); err != nil {
		log.Fatalf("Database dump failed: %v", err)
	}

	fmt.Println("Backup successful, sending email...")

	if err := sendEmail(); err != nil {
		log.Fatalf("Email sending failed: %v", err)
	}

	fmt.Println("Backup sent successfully!")
}
