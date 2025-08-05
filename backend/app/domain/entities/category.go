package entities

// TODO: Refer to the database schema. Now this is just sample.

// CREATE TABLE categories (
//     category_id VARCHAR(255) PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     display_order INTEGER NOT NULL DEFAULT 0,
//     image_url VARCHAR(255)
// );

type Category struct {
	CategoryID   string `json:"category_id,omitempty"`
	Name         string `json:"name,omitempty"`
	DisplayOrder int    `json:"display_order,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
}

func (c Category) GetTableName() string {
	return "categories"
}

// TODO: Implement validation methods for values
