Use for EKS pod via IRSA - update account number

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["secretsmanager:GetSecretValue"],
      "Resource": "arn:aws:secretsmanager:us-west-2:123456789012:secret:openai/api-key-*"
    }
  ]
}
```