# Prompts 

Spanish version: [esp/prompts.md](esp/prompts.md)

It would be stupid to do all the work without AI assistance today. So what was used.

## AI assistance

As AI assistance, I used Qwen 2.5 7b model, that I run on my PC. It's not a heavy one, and can be easily run on any modern PC with at least 16GB of RAM, [read](https://github.com/jzethar/Useful-Containers/blob/ollama/ollama/README.md)

Why self-hosted? Because I want to have full control over my data. I don't want to rely on third-party services that might not be transparent or secure. Plus, it's more cost-effective in the long run. 

Basically it helps autocompleting my code and comments. There is also an option to ask questions and get answers.

All this possible with Continue extension in VSCode. Just run the container and connect to it on port.

## Another AIs used

For this project, I also used group of AI to check and recheck the code that they generated. They are:
- ChatGPT 5
- Grok
- Gemini

Mostly was used ChatGPT 5 and Grok. 

### Collaboration

So, normally AI can be used for:
1. **Code generation**: I predefine the interface that I will have to implement, then I ask AI to generate the code.
2. **Testing**: I provide a piece of code and ask AI to write tests for it. Helps to automatize testing process.
3. **Code refactoring**: I provide a piece of code and ask AI to refactor it. In situation when we run small code pieces it doesn't used as much. But I remember once it was the situation when the linter had an error of the too deep code. So AI helped me to separate the code into smaller parts.
4. **Translation**: Living the life speaking 4 languages everyday there is a need of translation sometimes. So I can ask the AI to translate it for me using the context.

## Setup & Prompts

### Grok

My Grok is always set with this prompt:

```
Groks answer should be compact without deep diving until it will be asked for this
```

The big problem of Grok that it creates lots of unnecessary text that I'm wasting my time reading. So this prompt keeps it strict.

### ChatGPT 5

For ChatGPT 5, there is no necessary setup such prompt, cause it's smart enough to make answers short. But, if it switches to ChatGPT 4, then the same prompt will be used for it as well. 

For example, for generating the code I use this:

```
Can you generate a Go implementation for the following interface?
```

The same for tests. 

Normally it generates great answers and saves my time. But not all the time, so there should be a balance between generating and writing by myself.

## Final Notes

The best solution is a personal selfhosted assistant. If it can't be done, then I will use ChatGPT 5 or Grok as a fallback