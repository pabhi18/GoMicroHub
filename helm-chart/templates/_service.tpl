{{- define "common.service" -}}
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
    - protocol: TCP
      port: {{ .port }}
      targetPort: {{ .port }}
      name: main-port
    {{- if .additionalPorts }}
    {{- range .additionalPorts }}
    - protocol: TCP
      port: {{ .port }}
      targetPort: {{ .port }}
      name: {{ .name }}
    {{- end }}
    {{- end }}
{{- end -}}
