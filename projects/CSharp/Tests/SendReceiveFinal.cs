﻿using System;
using Xunit;

namespace Tests
{
    public class MyFinalSender : FBE.protoex.FinalSender, FBE.protoex.ISenderListener
    {
        public long OnSend(byte[] buffer, long offset, long size)
        {
            // Send nothing...
            return 0;
        }
    }

    public class MyFinalReceiver : FBE.protoex.FinalReceiver, FBE.protoex.IFinalReceiverListener
    {
        private bool _order;
        private bool _balance;
        private bool _account;

        public bool Check() { return _order && _balance && _account; }

        public void OnReceive(protoex.OrderMessage value) { _order = true; }
        public void OnReceive(protoex.BalanceMessage value) { _balance = true; }
        public void OnReceive(protoex.AccountMessage value) { _account = true; }
    }

    public class SendReceiveFinal
    {
        private static bool SendAndReceiveFinal(long index1, long index2)
        {
            var sender = new MyFinalSender();

            // Create and send a new order
            var order = new protoex.Order(1, "EURUSD", protoex.OrderSide.buy, protoex.OrderType.market, 1.23456, 1000.0, 0.0, 0.0);
            sender.Send(new protoex.OrderMessage(order));

            // Create and send a new balance wallet
            var balance = new protoex.Balance(new proto.Balance("USD", 1000.0), 100.0);
            sender.Send(new protoex.BalanceMessage(balance));

            // Create and send a new account with some orders
            var account = protoex.Account.Default;
            account.id = 1;
            account.name = "Test";
            account.state = protoex.StateEx.good;
            account.wallet.parent.currency = "USD";
            account.wallet.parent.amount = 1000.0;
            account.asset = new protoex.Balance(new proto.Balance("EUR", 100.0), 100.0);
            account.orders.Add(new protoex.Order(1, "EURUSD", protoex.OrderSide.buy, protoex.OrderType.market, 1.23456, 1000.0, 0.0, 0.0));
            account.orders.Add(new protoex.Order(2, "EURUSD", protoex.OrderSide.sell, protoex.OrderType.limit, 1.0, 100.0, 0.0, 0.0));
            account.orders.Add(new protoex.Order(3, "EURUSD", protoex.OrderSide.buy, protoex.OrderType.stop, 1.5, 10.0, 0.0, 0.0));
            sender.Send(new protoex.AccountMessage(account));

            var receiver = new MyFinalReceiver();

            // Receive data from the sender
            index1 %= sender.Buffer.Size;
            index2 %= sender.Buffer.Size;
            index2 = Math.Max(index1, index2);
            receiver.Receive(sender.Buffer.Data, 0, index1);
            receiver.Receive(sender.Buffer.Data, index1, index2 - index1);
            receiver.Receive(sender.Buffer.Data, index2, sender.Buffer.Size - index2);
            return receiver.Check();
        }

        [Fact(DisplayName = "Send & Receive (Final)")]
        public void SendAndReceiveFinalTest()
        {
            for (long i = 0; i < 1000; ++i)
                for (long j = 0; j < 1000; ++j)
                    Assert.True(SendAndReceiveFinal(i, j));
        }
    }
}
