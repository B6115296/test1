// Code generated by entc, DO NOT EDIT.

package groupofage

const (
	// Label holds the string label denoting the groupofage type in the database.
	Label = "group_of_age"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupOfAgeName holds the string denoting the group_of_age_name field in the database.
	FieldGroupOfAgeName = "group_of_age_name"
	// FieldGroupOfAgeAge holds the string denoting the group_of_age_age field in the database.
	FieldGroupOfAgeAge = "group_of_age_age"

	// EdgeGroupageProduct holds the string denoting the groupage_product edge name in mutations.
	EdgeGroupageProduct = "groupage_product"

	// Table holds the table name of the groupofage in the database.
	Table = "group_of_ages"
	// GroupageProductTable is the table the holds the groupage_product relation/edge.
	GroupageProductTable = "products"
	// GroupageProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	GroupageProductInverseTable = "products"
	// GroupageProductColumn is the table column denoting the groupage_product relation/edge.
	GroupageProductColumn = "group_of_age_id"
)

// Columns holds all SQL columns for groupofage fields.
var Columns = []string{
	FieldID,
	FieldGroupOfAgeName,
	FieldGroupOfAgeAge,
}
