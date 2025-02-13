{{- define "common.database" -}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .name }}
  labels:
    {{- include "helm-chart.labels" .root | nindent 4 }}
    app: {{ .name }}
spec:
  replicas: 1
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
        {{- if .resources }}
        resources:
          {{- toYaml .resources | nindent 12 }}
        {{- end }}
        {{- if .env }}
        env:
          {{- toYaml .env | nindent 10 }}
        {{- end }}
        ports:
          {{- if .ports }}
          {{- range .ports }}
          - containerPort: {{ .port }}
          {{- end }}
          {{- else }}
          - containerPort: {{ .port }}
          {{- end }}
---
apiVersion: v1
kind: Service
metadata:
  name: {{ .name }}
  labels:
    {{- include "helm-chart.labels" .root | nindent 4 }}
    app: {{ .name }}
spec:
  selector:
    app: {{ .name }}
  ports:
    {{- if .ports }}
    {{- range .ports }}
    - protocol: TCP
      name: {{ .name }}
      port: {{ .port }}
      targetPort: {{ .port }}
    {{- end }}
    {{- else }}
    - protocol: TCP
      port: {{ .port }}
      targetPort: {{ .port }}
    {{- end }}
  {{- if .clusterIP }}
  clusterIP: {{ .clusterIP }}
  {{- end }}
{{- end -}}