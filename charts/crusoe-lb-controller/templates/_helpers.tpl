{{- define "crusoe-lb-controller.name" -}}
{{ .Chart.Name | lower }}
{{- end -}}

{{- define "crusoe-lb-controller.fullname" -}}
{{- if .Release.Name -}}
{{ printf "%s-%s" .Release.Name (include "crusoe-lb-controller.name" .) | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{ include "crusoe-lb-controller.name" . | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}

{{- define "crusoe-lb-controller.serviceAccountName" -}}
{{ include "crusoe-lb-controller.fullname" . }}-sa
{{- end -}}

{{- define "crusoe-lb-controller.selectorLabels" -}}
app.kubernetes.io/name: {{ include "crusoe-lb-controller.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
