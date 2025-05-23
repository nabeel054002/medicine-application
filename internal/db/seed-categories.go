package db

import (
	"log"
	"fmt"
)

// seeds the database with default categories if they don't exist
func SeedCategories() error {
	categories := []string{
		"Pain Relief",
		"Antibiotics",
		"Vitamins",
		"Cold & Flu",
		"Digestive Health",
		"Skin Care",
	}

	for _, category := range categories {
		var count int
		err := DB.QueryRow("SELECT COUNT(*) FROM categories WHERE name = ?", category).Scan(&count)
		if err != nil {
			return err
		}

		if count == 0 {
			_, err := DB.Exec("INSERT INTO categories (name) VALUES (?)", category)
			if err != nil {
				return err
			}
			log.Printf("Category '%s' seeded successfully.", category)
		}
	}

	return nil
}

func SeedMedicines() error {
	medicines := []struct {
		ID           string
		Name         string
		CategoryName string
	}{
		{"med001", "Aspirin", "Pain Relief"},
		{"med002", "Amoxicillin", "Antibiotics"},
		{"med003", "Vitamin C", "Vitamins"},
		{"med004", "Paracetamol", "Pain Relief"},
		{"med005", "Coldrex", "Cold & Flu"},
		{"med006", "Lactulose", "Digestive Health"},
	}

	for _, medicine := range medicines {
		var categoryID int
		err := DB.QueryRow("SELECT id FROM categories WHERE name = ?", medicine.CategoryName).Scan(&categoryID)
		if err != nil {
			return fmt.Errorf("Error fetching category ID for '%s': %v", medicine.CategoryName, err)
		}

		var count int
		err = DB.QueryRow("SELECT COUNT(*) FROM medicines WHERE id = ?", medicine.ID).Scan(&count)
		if err != nil {
			return err
		}

		if count == 0 {
			_, err := DB.Exec(
				"INSERT INTO medicines (id, name, category_id) VALUES (?, ?, ?)",
				medicine.ID, medicine.Name, categoryID,
			)
			if err != nil {
				return fmt.Errorf("Error inserting medicine '%s': %v", medicine.Name, err)
			}
			log.Printf("Medicine '%s' with category '%s' and ID '%s' seeded successfully.",
				medicine.Name, medicine.CategoryName, medicine.ID)
		}
	}

	return nil
}
