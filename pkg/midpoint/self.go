package midpoint

import (
	"context"
	"log/slog"
	"time"
)

type SelfService struct {
	Service
}

func (s *SelfService) Read(ctx context.Context) (*Self, error) {

	self, err := s.client.Get[Self](ctx, "self?options=raw")
	if err != nil {
		slog.Error("error reading self", "error", err)
		return nil, err
	}

	return self, nil
}

type Self struct {
	User struct {
		Metadata struct {
			Ns      string `json:"@ns" yaml:"@ns"`
			Storage struct {
				CreateTimestamp time.Time `json:"createTimestamp" yaml:"createTimestamp"`
				CreateChannel   string    `json:"createChannel" yaml:"createChannel"`
			} `json:"storage" yaml:"storage"`
			Process struct {
				RequestTimestamp time.Time `json:"requestTimestamp" yaml:"requestTimestamp"`
			} `json:"process" yaml:"process"`
			ID int `json:"@id" yaml:"@id"`
		} `json:"@metadata" yaml:"@metadata"`
		Oid            string `json:"oid" yaml:"oid"`
		Version        string `json:"version" yaml:"version"`
		Name           string `json:"name" yaml:"name"`
		Indestructible bool   `json:"indestructible" yaml:"indestructible"`
		Assignment     []struct {
			Ns       string `json:"@ns" yaml:"@ns"`
			Metadata struct {
				Storage struct {
					CreateTimestamp time.Time `json:"createTimestamp" yaml:"createTimestamp"`
					CreateChannel   string    `json:"createChannel" yaml:"createChannel"`
				} `json:"storage" yaml:"storage"`
				Process struct {
					RequestTimestamp time.Time `json:"requestTimestamp" yaml:"requestTimestamp"`
				} `json:"process" yaml:"process"`
				Provenance struct {
					Acquisition struct {
						Channel   string    `json:"channel" yaml:"channel"`
						Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
						ID        string    `json:"@id" yaml:"@id"`
					} `json:"acquisition" yaml:"acquisition"`
				} `json:"provenance" yaml:"provenance"`
				ID int `json:"@id" yaml:"@id"`
			} `json:"@metadata" yaml:"@metadata"`
			ID         int    `json:"@id" yaml:"@id"`
			Identifier string `json:"identifier" yaml:"identifier"`
			TargetRef  struct {
				Oid      string `json:"oid" yaml:"oid"`
				Relation string `json:"relation" yaml:"relation"`
				Type     string `json:"type" yaml:"type"`
			} `json:"targetRef" yaml:"targetRef"`
			Activation struct {
				EffectiveStatus string `json:"effectiveStatus" yaml:"effectiveStatus"`
			} `json:"activation" yaml:"activation"`
		} `json:"assignment" yaml:"assignment"`
		Iteration      int    `json:"iteration" yaml:"iteration"`
		IterationToken string `json:"iterationToken" yaml:"iterationToken"`
		ArchetypeRef   struct {
			Oid      string `json:"oid" yaml:"oid"`
			Relation string `json:"relation" yaml:"relation"`
			Type     string `json:"type" yaml:"type"`
		} `json:"archetypeRef" yaml:"archetypeRef"`
		RoleMembershipRef []struct {
			Metadata struct {
				Ns      string `json:"@ns" yaml:"@ns"`
				ID      int    `json:"@id" yaml:"@id"`
				Storage struct {
					CreateTimestamp time.Time `json:"createTimestamp" yaml:"createTimestamp"`
				} `json:"storage" yaml:"storage"`
				Provenance struct {
					AssignmentPath struct {
						SourceRef struct {
							Oid      string `json:"oid" yaml:"oid"`
							Relation string `json:"relation" yaml:"relation"`
							Type     string `json:"type" yaml:"type"`
						} `json:"sourceRef" yaml:"sourceRef"`
						Segment struct {
							ID           int `json:"@id" yaml:"@id"`
							SegmentOrder int `json:"segmentOrder" yaml:"segmentOrder"`
							AssignmentID int `json:"assignmentId" yaml:"assignmentId"`
							TargetRef    struct {
								Oid      string `json:"oid" yaml:"oid"`
								Relation string `json:"relation" yaml:"relation"`
								Type     string `json:"type" yaml:"type"`
							} `json:"targetRef" yaml:"targetRef"`
							MatchingOrder bool `json:"matchingOrder" yaml:"matchingOrder"`
						} `json:"segment" yaml:"segment"`
					} `json:"assignmentPath" yaml:"assignmentPath"`
				} `json:"provenance" yaml:"provenance"`
			} `json:"@metadata" yaml:"@metadata"`
			Oid      string `json:"oid" yaml:"oid"`
			Relation string `json:"relation" yaml:"relation"`
			Type     string `json:"type" yaml:"type"`
		} `json:"roleMembershipRef" yaml:"roleMembershipRef"`
		Activation struct {
			AdministrativeStatus string    `json:"administrativeStatus" yaml:"administrativeStatus"`
			EffectiveStatus      string    `json:"effectiveStatus" yaml:"effectiveStatus"`
			EnableTimestamp      time.Time `json:"enableTimestamp" yaml:"enableTimestamp"`
			LockoutStatus        string    `json:"lockoutStatus" yaml:"lockoutStatus"`
		} `json:"activation" yaml:"activation"`
		Credentials struct {
			Password struct {
				Metadata struct {
					Ns      string `json:"@ns" yaml:"@ns"`
					Storage struct {
						CreateTimestamp time.Time `json:"createTimestamp" yaml:"createTimestamp"`
						CreateChannel   string    `json:"createChannel" yaml:"createChannel"`
					} `json:"storage" yaml:"storage"`
					ID int `json:"@id" yaml:"@id"`
				} `json:"@metadata" yaml:"@metadata"`
				Value struct {
					EncryptedData struct {
						EncryptionMethod struct {
							Algorithm string `json:"algorithm" yaml:"algorithm"`
						} `json:"encryptionMethod" yaml:"encryptionMethod"`
						KeyInfo struct {
							KeyName string `json:"keyName" yaml:"keyName"`
						} `json:"keyInfo" yaml:"keyInfo"`
						CipherData struct {
							CipherValue string `json:"cipherValue" yaml:"cipherValue"`
						} `json:"cipherData" yaml:"cipherData"`
					} `json:"encryptedData" yaml:"encryptedData"`
				} `json:"value" yaml:"value"`
			} `json:"password" yaml:"password"`
		} `json:"credentials" yaml:"credentials"`
		Behavior struct {
			Authentication []struct {
				Ns                  string `json:"@ns" yaml:"@ns"`
				ID                  int    `json:"@id" yaml:"@id"`
				LastSuccessfulLogin struct {
					Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
					From      string    `json:"from" yaml:"from"`
				} `json:"lastSuccessfulLogin" yaml:"lastSuccessfulLogin"`
				PreviousSuccessfulLogin struct {
					Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
					From      string    `json:"from" yaml:"from"`
				} `json:"previousSuccessfulLogin,omitempty" yaml:"previousSuccessfulLogin,omitempty"`
				AuthenticationAttempt struct {
					SequenceIdentifier           string `json:"sequenceIdentifier" yaml:"sequenceIdentifier"`
					ModuleIdentifier             string `json:"moduleIdentifier" yaml:"moduleIdentifier"`
					LastSuccessfulAuthentication struct {
						Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
						From      string    `json:"from" yaml:"from"`
					} `json:"lastSuccessfulAuthentication" yaml:"lastSuccessfulAuthentication"`
				} `json:"authenticationAttempt" yaml:"authenticationAttempt"`
				SequenceIdentifier string `json:"sequenceIdentifier" yaml:"sequenceIdentifier"`
			} `json:"authentication" yaml:"authentication"`
		} `json:"behavior" yaml:"behavior"`
		FullName   string `json:"fullName" yaml:"fullName"`
		GivenName  string `json:"givenName" yaml:"givenName"`
		FamilyName string `json:"familyName" yaml:"familyName"`
	} `json:"user" yaml:"user"`
}
