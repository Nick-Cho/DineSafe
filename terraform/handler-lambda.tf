resource "aws_iam_role" "handler_lambda_exec" {
  name = "handler-lambda"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "handler_lambda_policy" {
  role       = aws_iam_role.handler_lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "handler" {
  function_name = "handler"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.handler.key

  runtime = "go1.x"
  handler = "function.handler"

  source_code_hash = data.archive_file.handler.output_base64sha256

  role = aws_iam_role.handler_lambda_exec.arn
}


resource "aws_cloudwatch_log_group" "handler_lambda" {
  name = "/aws/lambda/${aws_lambda_function.handler.function_name}"
}

data "archive_file" "handler" {
  type        = "zip"
  source_dir  = "../cmd/routes/getRestaurantReviews/"
  output_path = "../cmd/routes/getRestaurantReviews/main.zip"
}

resource "aws_s3_object" "handler" {
  bucket  = aws_s3_bucket.lambda_bucket.id
  key    = "main.zip"
  source = data.archive_file.handler.output_path
  etag = filemd5(data.archive_file.handler.output_path)
}