package task

const TASK = `apiVersion: batch/v1
kind: Job
metadata:
  name: {{.RELEASE_NAME}}-console-{{.UUID}}
spec:
  ttlSecondsAfterFinished: 5
  template:
    metadata:
      name: {{.RELEASE_NAME}}-console-{{.UUID}}
    spec:
      containers:
        - name: {{.RELEASE_NAME}}-console
          image: {{.IMAGE_URL}}
          resources:
            requests:
              memory: "1Gi"
          imagePullPolicy: IfNotPresent
          command:
          - /bin/herokuish
          args:
          - procfile
          - exec
          - {{.COMMAND}}
          tty: true
          stdin: true
          envFrom:
          - configMapRef:
              name: {{.RELEASE_NAME}}-config-map
          - secretRef:
              name: {{.RELEASE_NAME}}-secret
      restartPolicy: "Never"`