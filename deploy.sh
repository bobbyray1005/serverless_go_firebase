gcloud functions deploy gofirebase --env-vars-file .env.yml  \
--runtime go113 --trigger-http --allow-unauthenticated --entry-point User \
--memory 256MB --region europe-west3 --max-instances 1 \
--source=.