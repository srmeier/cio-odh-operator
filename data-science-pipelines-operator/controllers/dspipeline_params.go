/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"fmt"
	mf "github.com/manifestival/manifestival"
	dspipelinesiov1alpha1 "github.com/opendatahub-io/data-science-pipelines-operator/api/v1alpha1"
	"math/rand"
	"time"
)

const (
	defaultDBHostPrefix                  = "mariadb"
	defaultDBHostPort                    = "3306"
	defaultMinioHostPrefix               = "minio"
	defaultMinioPort                     = "9000"
	defaultArtifactScriptConfigMapPrefix = "ds-pipeline-artifact-script-"
	defaultArtifactScriptConfigMapKey    = "artifact_script"
	defaultDSPServicePrefix              = "ds-pipeline"
)

type DSPipelineParams struct {
	Name                             string
	Namespace                        string
	Owner                            mf.Owner
	APIServerImage                   string
	PipelineRuntime                  string
	StripEOF                         string
	ArtifactScript                   string
	ArtifactImage                    string
	InjectDefaultScript              string
	ApplyTektonCustomResource        string
	TerminateStatus                  string
	AutoUpdatePipelineDefaultVersion string
	DBConfigCONMAXLifetimeSec        string
	ObjectStoreConfigBucketName      string
	S3CredentialsSecretName          string
	AccessKeySecretKey               string
	SecretKey                        string
	ObjectStoreConfigSecure          string
	MinioServiceServiceHost          string
	MinioServiceServicePort          string
	ArtifactBucket                   string
	ArtifactEndpoint                 string
	ArtifactEndpointScheme           string
	ArtifactLogs                     string
	ArtifactTracking                 string
	ArtifactScriptConfigMapName      string
	ArtifactScriptConfigMapKey       string
	MariaDBImage                     string
	DBUser                           string
	DBPasswordSecretKey              string
	DBPasswordSecret                 string
	DBName                           string
	DBHost                           string
	DBPort                           string
	DBPassword                       string
	SecretKeyVal                     string
	AccessKey                        string
	MinioImage                       string
	DBServiceName                    string
	MinioServiceName                 string
	PersistenceAgentImage            string
	APIServerServiceName             string
	ScheduledWorkflowImage           string
	VisualizationServerImage         string
	ViewerCrdImage                   string
	CronScheduleTimezone             string
	MLPipelineUIImage                string
	MaxNumViewer                     string
}

func passwordGen(n int) string {
	rand.Seed(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func (r *DSPipelineParams) UsingCustomDB(dsp *dspipelinesiov1alpha1.DSPipeline) (bool, error) {
	if dsp.Spec.Database.CustomDB != (dspipelinesiov1alpha1.CustomDB{}) {
		return true, nil
	} else if dsp.Spec.Database.MariaDB != (dspipelinesiov1alpha1.MariaDB{}) {
		return false, nil
	}
	return false, fmt.Errorf("no Database specified for DS-Pipeline resource")
}

func (r *DSPipelineParams) UsingCustomStorage(dsp *dspipelinesiov1alpha1.DSPipeline) (bool, error) {
	if dsp.Spec.Storage.CustomStorage != (dspipelinesiov1alpha1.CustomStorage{}) {
		return true, nil
	} else if dsp.Spec.Storage.Minio != (dspipelinesiov1alpha1.Minio{}) {
		return false, nil
	}
	return false, fmt.Errorf("no Database specified for DS-Pipeline resource")
}

func (r *DSPipelineParams) ExtractParams(dsp *dspipelinesiov1alpha1.DSPipeline) error {
	r.Name = dsp.Name
	r.Namespace = dsp.Namespace
	r.Owner = dsp
	r.APIServerImage = dsp.Spec.APIServer.Image
	r.PipelineRuntime = "tekton"
	r.StripEOF = "true"
	r.ObjectStoreConfigSecure = "false"
	r.ArtifactEndpointScheme = "http:/"
	r.ArtifactLogs = "false"
	r.ArtifactTracking = "true"
	r.ArtifactImage = dsp.Spec.APIServer.ArtifactImage
	r.InjectDefaultScript = "true"
	r.ApplyTektonCustomResource = "true"
	r.TerminateStatus = "Cancelled"
	r.AutoUpdatePipelineDefaultVersion = "true"
	r.DBConfigCONMAXLifetimeSec = "120"
	r.PersistenceAgentImage = dsp.Spec.PersistentAgent.Image
	r.APIServerServiceName = fmt.Sprintf("%s-%s", defaultDSPServicePrefix, r.Name)
	r.ScheduledWorkflowImage = dsp.Spec.ScheduledWorkflow.Image
	r.ViewerCrdImage = dsp.Spec.ViewerCRD.Image
	r.CronScheduleTimezone = "UTC"
	r.MLPipelineUIImage = dsp.Spec.MlPipelineUI.Image
	r.MaxNumViewer = "50"

	if dsp.Spec.APIServer.ArtifactScriptConfigMap != (dspipelinesiov1alpha1.ArtifactScriptConfigMap{}) {
		r.ArtifactScriptConfigMapName = dsp.Spec.APIServer.ArtifactScriptConfigMap.Name
		r.ArtifactScriptConfigMapKey = dsp.Spec.APIServer.ArtifactScriptConfigMap.Key
	} else {
		r.ArtifactScriptConfigMapName = defaultArtifactScriptConfigMapPrefix + dsp.Name
		r.ArtifactScriptConfigMapKey = defaultArtifactScriptConfigMapKey
	}

	usingCustomDB, err := r.UsingCustomDB(dsp)
	if err != nil {
		return err
	}

	if usingCustomDB {
		customDB := dsp.Spec.Database.CustomDB
		r.DBUser = customDB.Username
		r.DBName = customDB.DBName
		r.DBPasswordSecretKey = customDB.PasswordSecret.Key
		r.DBPasswordSecret = customDB.PasswordSecret.Name
		r.DBHost = customDB.Host
		r.DBPort = customDB.Port
	} else {
		mariaDB := dsp.Spec.Database.MariaDB
		r.MariaDBImage = mariaDB.Image
		r.DBUser = mariaDB.Username
		r.DBName = mariaDB.DBName
		r.DBPasswordSecretKey = mariaDB.PasswordSecret.Key
		r.DBPasswordSecret = mariaDB.PasswordSecret.Name
		r.DBPassword = passwordGen(12)
		r.DBServiceName = defaultDBHostPrefix + "-" + r.Name
		r.DBHost = fmt.Sprintf("%s.%s.svc.cluster.local", r.DBServiceName, r.Namespace)
		r.DBPort = defaultDBHostPort
	}

	usingCustomStorage, err := r.UsingCustomStorage(dsp)
	if err != nil {
		return err
	}

	if usingCustomStorage {
		storage := dsp.Spec.Storage.CustomStorage
		r.ObjectStoreConfigBucketName = storage.Bucket
		r.S3CredentialsSecretName = storage.SecretName
		r.AccessKeySecretKey = storage.AccessKey
		r.SecretKey = storage.SecretKey
		r.MinioServiceServiceHost = storage.Host
		r.MinioServiceServicePort = storage.Port
		r.ArtifactBucket = storage.Bucket
		r.ArtifactEndpoint = storage.Host + ":" + storage.Port
	} else {
		storage := dsp.Spec.Storage.Minio
		r.ObjectStoreConfigBucketName = storage.Bucket
		r.S3CredentialsSecretName = storage.SecretName
		r.AccessKeySecretKey = storage.AccessKey
		r.SecretKey = storage.SecretKey
		r.MinioServiceName = defaultMinioHostPrefix + "-" + r.Name
		r.MinioServiceServiceHost = fmt.Sprintf("%s.%s.svc.cluster.local", r.MinioServiceName, r.Namespace)
		r.MinioServiceServicePort = defaultMinioPort
		r.MinioImage = dsp.Spec.Minio.Image
		r.ArtifactBucket = storage.Bucket
		r.ArtifactEndpoint = defaultMinioHostPrefix + "-" + r.Name + ":" + defaultMinioPort
		r.SecretKeyVal = passwordGen(12)
		r.AccessKey = passwordGen(12)
	}

	return nil
}
