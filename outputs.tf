output "vpc_id" {
  value = aws_vpc.demo_vpc.id
}

output "sg_id" {
  value = aws_security_group.demo_sg.id
}

output "ec2_id" {
  value = aws_instance.demo_ec2.id
}

output "bucket_name" {
  value = aws_s3_bucket.demo.id
}
