using System.Net.Mime;
using System.Text;
using System.Text.Json;
using DeadShotDotnetLib.Interfaces;
using DeadShotDotnetLib.Models;

namespace DeadShotDotnetLib;

public class DeadShotDotnetLib(DeadShotDotnetLibConfig config) : IDeadShotDotnetLib
{
    public async Task<bool> SendAsync(LogModel log)
    {
        try
        {
            HttpClient hClient = new();
            string url = $"{config.DeadShotHost}:{config.DeadShotPort}";
            HttpRequestMessage req = new(HttpMethod.Post, url);
            string body = JsonSerializer.Serialize(log);
            HttpContent content = new StringContent(body, Encoding.UTF8, MediaTypeNames.Application.Json);
            req.Content = content;
            HttpResponseMessage res = await hClient.SendAsync(req);
            res.EnsureSuccessStatusCode();
            return true;
        }
        catch (Exception e)
        {
            Console.WriteLine(e.Message);
            Console.WriteLine(e.StackTrace);
            return false;
        }
    }

    public bool Send(LogModel log)
    {
        try
        {
            HttpClient hClient = new();
            string url = $"{config.DeadShotHost}:{config.DeadShotPort}";
            HttpRequestMessage req = new(HttpMethod.Post, url);
            string body = JsonSerializer.Serialize(log);
            HttpContent content = new StringContent(body, Encoding.UTF8, MediaTypeNames.Application.Json);
            req.Content = content;
            HttpResponseMessage res = hClient.Send(req);
            res.EnsureSuccessStatusCode();
            return true;
        }
        catch (Exception e)
        {
            Console.WriteLine(e.Message);
            Console.WriteLine(e.StackTrace);
            return false;
        }
    }
}