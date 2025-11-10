resource "aws_instance" "demo_ec2" {
  ami           = "ami-12345678"  # đổi thành AMI hợp lệ hoặc LocalStack AMI
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.public.id

  # Sửa tên Security Group đúng với resource hiện tại
  vpc_security_group_ids = [aws_security_group.demo_sg.id]

  tags = {
    Name = "demo-ec2"
  }
}
