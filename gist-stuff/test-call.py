from openai import OpenAI
client = OpenAI()

response = client.responses.create(
    model="gpt-4o",
    input="Write a humorous devops haiku."
)

print(response.output_text)