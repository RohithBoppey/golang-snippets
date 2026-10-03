When using Redis pub sub, it is like a radio -- only live pub and sub works, and not storage 

so `dbsize` and `keys *` will never give you this entries in that 

but you can go into the cli, and do this: 
`subscribe topic_name` 

and then trigger from the api, and you can see the messages coming in live! 