resource "aws_s3_bucket" "haikube_bucket" {
  bucket        = "haikube"
  force_destroy = true

  tags = {
    Name = "Terraform State Bucket"
  }
}

resource "aws_s3_bucket_versioning" "haikube_bucket_versioning" {
  bucket = aws_s3_bucket.haikube_bucket.id

  versioning_configuration {
    status = "Enabled"
  }
}

output "bucket" {
  value = aws_s3_bucket.haikube_bucket.bucket
}