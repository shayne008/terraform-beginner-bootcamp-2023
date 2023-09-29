variable "user_uuid" {
  description = "User UUID"
  type        = string
  validation {
    condition     = can(regex("^([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})$", var.user_uuid))
    error_message = "Invalid user UUID format. It should be in the format of a UUID (e.g., 123e4567-e89b-12d3-a456-426655440000)."
  }
}

variable "bucket_name" {
  type        = string
  description = "Name of the AWS S3 bucket"

  validation {
    condition     = can(regex("^[a-z0-9.-]{3,63}$", var.bucket_name))
    error_message = "Bucket name must be between 3 and 63 characters, lowercase, and can only contain lowercase letters, numbers, hyphens, and periods."
  }
}

variable "index_html_filepath" {
  description = "Path to the index.html file"
  type        = string

  validation {
    condition     = fileexists(var.index_html_filepath)
    error_message = "The specified index_html_filepath is not a valid file path."
  }
}

variable "error_html_filepath" {
  description = "Path to the error.html file"
  type        = string

  validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The specified error_html_filepath is not a valid file path."
  }
}

variable "content_version" {
  description = "The content version, a positive integer starting at 1."
  type        = number
  validation {
    condition     = var.content_version >= 1 
    error_message = "Content version must be a positive integer starting at 1."
  }
}

