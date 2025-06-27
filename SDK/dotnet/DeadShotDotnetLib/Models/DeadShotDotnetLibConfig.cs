namespace DeadShotDotnetLib.Models;

public record DeadShotDotnetLibConfig
{
    public string DeadShotHost { get; set; }
    public int DeadShotPort { get; set; }
}