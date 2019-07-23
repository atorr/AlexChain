import json
import hashlib

class Wallet(object):
  def __init__(self):
    self.pvt = ""
    self.pub = ""
    self.addr = ""
    self.amount = 0.0
    self.history = []

  def tx_send(self, recipient, amount):
    tx = {
      "sender": self.pub,
      "recipient": recipient,
      "amount": amount
    }
    #TODO: broadcast transaction in JSON    

    #TODO: remove amount from total funds of the current wallet
    self.amount -= amount
    self.history.append(tx)
    
    pass

  def tx_receive(self, sender, amount):
    self.amount += amount
  
  def show_details(self, my_pub):
    wallet = {
      "address": self.addr,
      "amount": self.amount,
      "history": self.history
    }
    return json.loads(wallet)

  
