# `sanitise-object-name`

This Task can be used to sanitise a string that will be used for naming a k8s object. This task converts the string to a value that can be used for an object name. It converts any upercase character to lowercase, and converts and non-alphanumeric character that is not `-` or `.` to `-` and then trims these characters from either side of the string

## Parameters
* **input-string**: The string to be sanitised

## Workspaces

## Results
* **output-string**: Sanitised output string
