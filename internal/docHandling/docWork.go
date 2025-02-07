package docHandling

import "fmt"

// DocView is a function that allows a user to view a document
func DocView(docID string) error {
	// Get the document from the database
	// Display the document to the user
	fmt.Println("DocView", docID)
	return nil
}

// DocUpload is a function that allows a user to upload a document
func DocUpload(docID string) error {
	// Get the document from the user
	// Upload the document to the database
	fmt.Println("DocUpload", docID)
	return nil
}

// DocRemove is a function that allows a user to remove a document
func DocRemove(docID string) error {
	// Get the document from the user
	// Remove the document from the database
	fmt.Println("DocRemove", docID)
	return nil
}
