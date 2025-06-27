using DeadShotDotnetLib.Interfaces;
using DeadShotDotnetLib.Models;
using Microsoft.Extensions.DependencyInjection;

namespace DeadShotDotnetLib;

public static class DependencyInjection
{
    public static IServiceCollection AddDeadShot(this IServiceCollection services, DeadShotDotnetLibConfig config)
    {
        services.AddSingleton(config);
        services.AddTransient<IDeadShotDotnetLib, DeadShotDotnetLib>();

        return services;
    }
}