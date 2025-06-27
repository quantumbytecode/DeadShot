using System.Text.Json.Serialization;

namespace DeadShotDotnetLib.Models;

public record LogModel
{
    [JsonPropertyName("id")]
    public int ID { get; set; }
    [JsonPropertyName("method")]
    public string Method { get; set; }
    [JsonPropertyName("url")]
    public string URL { get; set; }
    [JsonPropertyName("headers")]
    public string Headers { get; set; }
    [JsonPropertyName("query_params")]
    public string QueryParams { get; set; }
    [JsonPropertyName("body")]
    public string Body { get; set; }
    [JsonPropertyName("received_at")]
    public DateTime ReceivedAt { get; set; }
    [JsonPropertyName("status_code")]
    public int StatusCode { get; set; }
    [JsonPropertyName("response_headers")]
    public string ResponseHeaders { get; set; }
    [JsonPropertyName("response_body")]
    public string ResponseBody { get; set; }
    [JsonPropertyName("tags")]
    public string Tags { get; set; }
    [JsonPropertyName("source")]
    public string Source { get; set; }
    [JsonPropertyName("replayed")]
    public bool Replayed { get; set; }
    [JsonPropertyName("replayed_count")]
    public int ReplayedCount { get; set; }
    [JsonPropertyName("error")]
    public string Error { get; set; }
}