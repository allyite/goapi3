# goapi3


Fetches covid-stats for India through a publically available covid-stats-api and persists the data in mongo-db collection.

Takes user's Geo-coordinates as input & returns the covid-stats for the State (India only), in which those coordinates lie.
Swagger documentation: https://app.swaggerhub.com/apis-docs/allyite/goapi_covid/0.2

End points:
/update 
to fetch latest covid stats & update data in mongo collection

/stats   ?lat=""&long=""
to return the covid stats for the state, based on user's geo-coordinates (latitude & longitude)
