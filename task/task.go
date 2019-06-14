package task

const TASK = `apiVersion: batch/v1
kind: Job
metadata:
  name: {{.release_name}}-console-{{.unique_id}}
spec:
  ttlSecondsAfterFinished: 5
  template:
    metadata:
      name: {{.release_name}}-console-{{.unique_id}}
    spec:
      containers:
        - name: {{.release_name}}-console
          image: {{.image_url}}
          imagePullPolicy: IfNotPresent
          command:
          - /bin/herokuish
          args:
          - procfile
          - exec
          - {{.command}}
          tty: true
          stdin: true
          envFrom:
          - configMapRef:
              name: {{.release_name}}-config-map
          - secretRef:
              name: {{.release_name}}-secret
      restartPolicy: "Never"`