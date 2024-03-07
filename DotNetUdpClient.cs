namespace DotNetUDPClient
{
    using System;
    using System.Net;
    using System.Net.Sockets;
    using System.Text;
    using System.Threading.Tasks;


    class Program
    {
        static async Task Main(string[] args)
        {
            UdpClient client = new UdpClient();
            IPEndPoint serverEndPoint = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 8080);

            int count = 0;
            while (true)
            {
                string message = $"Hi server, {count}";
                byte[] data = Encoding.ASCII.GetBytes(message);
                await client.SendAsync(data, data.Length, serverEndPoint);
                Console.WriteLine($"Sent: {message} to {serverEndPoint}");

                // Receive response from the server
                UdpReceiveResult result = await client.ReceiveAsync();
                string responseMessage = Encoding.ASCII.GetString(result.Buffer);
                Console.WriteLine($"Received from server: {serverEndPoint}, message: {responseMessage}");

                count += 1;

                //System.Threading.Thread.Sleep(100);
                await Task.Delay(100); // Wait for 0.1 second before sending the next message
            }
        }
    }
}