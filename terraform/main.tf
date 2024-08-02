terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-west-2"
  access_key = "***"
  secret_key = "***"
}

resource "aws_vpc" "tt_vpc" {
  cidr_block = "10.100.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = "tt-vpc"
  }
}

resource "aws_internet_gateway" "tt_internet_gateway" {
  vpc_id = aws_vpc.tt_vpc.id

  tags = {
    Name = "tt-internet-gateway"
  }
}

resource "aws_subnet" "tt_public_subnet" {
  vpc_id            = aws_vpc.tt_vpc.id
  cidr_block        = "10.100.10.0/24"
  availability_zone = "us-west-2a"
  map_public_ip_on_launch = true

  tags = {
    Name = "tt-subnet-public"
  }
}

resource "aws_subnet" "tt_private_subnet" {
  vpc_id            = aws_vpc.tt_vpc.id
  cidr_block        = "10.100.20.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "tt-subnet-private"
  }
}

resource "aws_route_table" "tt_public_route_table" {
  vpc_id = aws_vpc.tt_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.tt_internet_gateway.id
  }

  tags = {
    Name = "tt-public-route-table"
  }
}

resource "aws_route_table_association" "tt_pub_subnet_to_pub_rt_association" {
  subnet_id      = aws_subnet.tt_public_subnet.id
  route_table_id = aws_route_table.tt_public_route_table.id
}

# data "aws_ami_ids" "aws_ubuntu" {
#   executable_users = ["self"]
#   owners      = ["self"]

#  filter {
#     name   = "name"
#     values = ["al2023-ami-.*-kernel-6.1-x86_64"]
#   }
# }

resource "aws_instance" "aws_ubuntu_nginx_private" {
  instance_type = "t2.micro"
  ami           = "ami-078701cc0905d44e4"
  # key_name               = var.key_name
  user_data = file("userdata.tpl")
  subnet_id = aws_subnet.tt_private_subnet.id
  tags = {
    Name = "tf-nginx-private"
  }
} 

resource "aws_instance" "aws_ubuntu_nginx_public" {
  instance_type = "t2.micro"
  ami           = "ami-078701cc0905d44e4"
  # key_name               = var.key_name
  user_data = file("userdata.tpl")
  subnet_id = aws_subnet.tt_public_subnet.id
  tags = {
    Name = "tf-nginx-public"
  }
} 
