SELECT 
    time, metric, value
FROM metrics
WHERE
    time >= :start AND 
    time <= :end
;
