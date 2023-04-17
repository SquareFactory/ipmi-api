# IPMI API

A golang ipmitool API for [Cluster Factory](https://github.com/SquareFactory/ClusterFactory-CE).

---

## Endpoints


| Method | Endpoints          | Description                                                       |
| -------- | -------------------- | ------------------------------------------------------------------- |
| POST   | /host/:host/on     | Power up chassis.                                                 |
| POST   | /host/:host/off    | Power down chassis. Does not initiate a clean shutdown of the OS. |
| POST   | /host/:host/cycle  | Power cycle                                                       |
| GET    | /host/:host/status | Show current power status.                                        |
| POST   | /host/:host/soft   | Soft-shutdown of OS via ACPI.                                     |
| POST   | /host/:host/reset  | Hard reset.                                                       |

:host being the BMC IP adress / hostname.
