
# PowerBall Exporter

An exporter for randomly generated seeds for powerball numbers

``` diff
                                                       /$$                 /$$ /$$
                                                      | $$                | $$| $$
  /$$$$$$   /$$$$$$  /$$  /$$  /$$  /$$$$$$   /$$$$$$ | $$$$$$$   /$$$$$$ | $$| $$
|/$$__  $$ /$$__  $$| $$ | $$ | $$ /$$__  $$ /$$__  $$| $$__  $$ |____  $$| $$| $$
| $$  \ $$| $$  \ $$| $$ | $$ | $$| $$$$$$$$| $$  \__/| $$  \ $$  /$$$$$$$| $$| $$
| $$  | $$| $$  | $$| $$ | $$ | $$| $$_____/| $$      | $$  | $$ /$$__  $$| $$| $$
| $$$$$$$/|  $$$$$$/|  $$$$$/$$$$/|  $$$$$$$| $$      | $$$$$$$/|  $$$$$$$| $$| $$
| $$____/  \______/  \_____/\___/  \_______/|__/      |_______/  \_______/|__/|__/
| $$                                                                              
| $$                                                                              
|__/                                                                             
```

## Run
Run docker-compose to run this exporter

### Example Image
![alt text](Images/powerball-metrics.png)

### Metric Examples

```
powerball_metric{ball1, ball2, ball3, ball4, ball5, powerball} X value where x is the count of random generated value. This is fairly useless but helps distinguish duplicate values posted.
```

```
powerball_metric{ball1="1",ball2="19",ball3="30",ball4="53",ball5="68",pb="18"} 2
```