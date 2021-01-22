FILE=config/config.yaml
if [ -f "$FILE" ]; then
  echo "$FILE exist"
else
  echo "$FILE does not exist"
  cp config/config.sample.yaml $FILE
fi