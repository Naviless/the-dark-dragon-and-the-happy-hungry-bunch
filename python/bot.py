import discord
from discord.ext import commands
import json

#async def send_message(message, user_message, is_private):
   # try:
   #     response = responses.handle_response(user_message)
      #  await message.author.send(response) if is_private else await message.channel.send(response)     
 #   except Exception as e:
  #      print(e)    

def run_discord_bot():
    
    config_file = open('config.json')
    
    config_data = json.load(config_file)
    
    Token = config_data['Token']
    
    intents = discord.Intents.all()
    
    description = "If this doesn't work I'll end it"
    
    bot = discord.Bot(command_prefix=commands.when_mentioned_or('!'),description=description, intents=intents)
    
    @bot.event
    async def on_ready():
        print(f"{bot.user} is now running!")
        
    
   # @bot.event
   # async def on_message(message):
    #    if message.author == bot.user:
     #       return
        
      #  username = str(message.author)
      #  user_message = str(message.content) 
      #  channel = str(message.channel)
        
     #   print(f"{username} said:'{user_message}' ({channel})")
        
    @bot.slash_command()
    async def ping(ctx):
        await ctx.respond(f"Pong! Latency is {bot.latency} seconds.")
        
    @bot.slash_command()
    async def calculate(ctx, number1, operator, number2):
        def add(num1, num2):
            return int(num1) + int(num2)
        def multiply(num1, num2):
            return int(num1) * int(num2)
        def substract(num1, num2):
            return int(num1) - int(num2)
        def divide(num1, num2):
            return int(num1) / int(num2)
        
        if operator == "+":
            result = add(number1, number2)
        
        elif operator == "-":
            result = substract(number1, number2)
            
        elif operator == "*":
            result = multiply(number1, number2)
            
        elif operator == "/":
            result = divide(number1, number2)
        
        await ctx.respond(f"{number1} {operator} {number2} = {result}")        
            
            
            
            
            
    bot.run(Token)
