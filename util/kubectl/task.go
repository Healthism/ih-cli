package kubectl

const JOB_TEMPLATE = `apiVersion: batch/v1
kind: Job
metadata:
  name: {{.Name}}
spec:
  ttlSecondsAfterFinished: 5
  template:
    metadata:
      name: {{.Name}}
    spec:
      containers:
        - name: {{.Release}}-console
          image: {{.Image}}
          resources:
            requests:
              memory: "1Gi"
          imagePullPolicy: IfNotPresent
          command:
          - /bin/herokuish
          args:
          - procfile
          - exec
          - {{.Command}}
          tty: true
          stdin: true
          envFrom:
          - configMapRef:
              name: {{.Release}}-config-map
          - secretRef:
              name: {{.Release}}-secret
      restartPolicy: "Never"`
