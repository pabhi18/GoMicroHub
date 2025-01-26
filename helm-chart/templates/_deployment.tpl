{{- define "common.deployment" -}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .name }}
  labels:
    {{- include "helm-chart.labels" .root | nindent 4 }}
    app: {{ .name }}
spec:
  replicas: {{ .replicas }}
  selector:
    matchLabels:
      app: {{ .name }}
  template:
    metadata:
      labels:
        app: {{ .name }}
    spec:
      containers:
      - name: {{ .name }}
        image: "{{ .image }}:{{ .tag }}"
        resources:
          {{- toYaml .resources | nindent 12 }}
        ports:
          - containerPort: {{ .port }}
          {{- if .additionalPorts }}
          {{- range .additionalPorts }}
          - containerPort: {{ .port }}
          {{- end }}
          {{- end }}
        {{- if .env }}
        env:
          {{- toYaml .env | nindent 10 }}
        {{- end }}
{{- end -}}