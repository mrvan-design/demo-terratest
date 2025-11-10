resource "aws_security_group" "demo_sg" {
  name        = "demo-sg"
  description = "Demo Security Group"
  vpc_id      = aws_vpc.demo_vpc.id  # dùng trực tiếp từ root module

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "demo-sg"
  }
}