#type 

---

# Type: [[schedule]]

## Step 1: Describe Purpose

contains a schedule that will be used to trigger notifications

## Step 2: DynamoDB Composite Key Structure

### [[Partition Key]] (PK)
- Key Name:  `pk`
- Key Format: `<office_id>::0`
	- `::0` - indicates the default [[write shard]]
- Example: `1234-5678-9101::0`

### [[Sort Key]] (SK)
- Key Name: `sk`
- Key Format: `<type>::<name>`
- Example:  `platform-team::weekly`

## Step 3: Define Fields

List all fields, whether required or optional. 
- `userID`: string
- `officeID`: string
- `name`: string
- `is_active`: bool
- `type`: string

## Step 4: Describe Fields

Explain each field's purpose.\
- `userID`: the ID of the [[user]] who made this schedule
- `officeID`: the ID of the office this schedule is for
- `name`: a name provided by the schedule's creator
- `is_active`: denotes if the schedule will request availability
- `type`:  a free-form flexible way to categorize schedules
## Step 5: Add Constraints

Any limitations or constraints on the fields.
- none
## Step 6: Define Relationships

How does this type relate to other types?
- `userID` links back to [[user]] who created it
- `officeID` - links to the office this is for

## Step 7: List Methods

Any methods or functions related to this type.
- `getOffice()`: Office
- `copyToOffice(officeID string)`: NewOffice

## Step 8: Describe Methods

Provide a description or use-case for each method.
- `getOffice`: fetches the office this schedule is linked to
- `copyToOffice`: copies this schedule to another office, but ensure `is_active` is false

## Step 9: Usage Scenarios

Where and how will this type be used?
- Creating a new schedule
- Running the [[office notifier]]
- copying a schedule from one office to another

## Step 10: Additional Notes

Any other information, constraints, or code snippets.

