resource "aws_apigatewayv2_integration" "get_restaurant_reviews_handler" {
  api_id = aws_apigatewayv2_api.main.id

  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.handler.invoke_arn
}

resource "aws_apigatewayv2_route" "get_restaurant_reviews_route" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /getRestaurantReviews/{street_address}"
    
  target = "integrations/${aws_apigatewayv2_integration.get_restaurant_reviews_handler.id}"
}

resource "aws_lambda_permission" "api_gw_get_restaurant_reviews" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.handler.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}