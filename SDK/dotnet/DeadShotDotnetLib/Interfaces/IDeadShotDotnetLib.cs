using DeadShotDotnetLib.Models;

namespace DeadShotDotnetLib.Interfaces;

public interface IDeadShotDotnetLib
{
    public bool Send(LogModel log);
}