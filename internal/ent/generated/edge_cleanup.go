// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/privacy"
	"github.com/rs/zerolog/log"
	"github.com/theopenlane/core/internal/ent/generated/contact"
	"github.com/theopenlane/core/internal/ent/generated/documentdata"
	"github.com/theopenlane/core/internal/ent/generated/emailverificationtoken"
	"github.com/theopenlane/core/internal/ent/generated/entitlement"
	"github.com/theopenlane/core/internal/ent/generated/entitlementplan"
	"github.com/theopenlane/core/internal/ent/generated/entitlementplanfeature"
	"github.com/theopenlane/core/internal/ent/generated/entity"
	"github.com/theopenlane/core/internal/ent/generated/entitytype"
	"github.com/theopenlane/core/internal/ent/generated/group"
	"github.com/theopenlane/core/internal/ent/generated/groupmembership"
	"github.com/theopenlane/core/internal/ent/generated/groupsetting"
	"github.com/theopenlane/core/internal/ent/generated/integration"
	"github.com/theopenlane/core/internal/ent/generated/invite"
	"github.com/theopenlane/core/internal/ent/generated/note"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/organizationsetting"
	"github.com/theopenlane/core/internal/ent/generated/orgmembership"
	"github.com/theopenlane/core/internal/ent/generated/passwordresettoken"
	"github.com/theopenlane/core/internal/ent/generated/personalaccesstoken"
	"github.com/theopenlane/core/internal/ent/generated/program"
	"github.com/theopenlane/core/internal/ent/generated/programmembership"
	"github.com/theopenlane/core/internal/ent/generated/subscriber"
	"github.com/theopenlane/core/internal/ent/generated/task"
	"github.com/theopenlane/core/internal/ent/generated/template"
	"github.com/theopenlane/core/internal/ent/generated/tfasetting"
	"github.com/theopenlane/core/internal/ent/generated/user"
	"github.com/theopenlane/core/internal/ent/generated/usersetting"
	"github.com/theopenlane/core/internal/ent/generated/webauthn"
	"github.com/theopenlane/core/internal/ent/generated/webhook"
)

func APITokenEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup apitoken edge"))

	return nil
}

func ActionPlanEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup actionplan edge"))

	return nil
}

func ActionPlanHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup actionplanhistory edge"))

	return nil
}

func ContactEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup contact edge"))

	return nil
}

func ContactHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup contacthistory edge"))

	return nil
}

func ControlEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup control edge"))

	return nil
}

func ControlHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup controlhistory edge"))

	return nil
}

func ControlObjectiveEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup controlobjective edge"))

	return nil
}

func ControlObjectiveHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup controlobjectivehistory edge"))

	return nil
}

func DocumentDataEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup documentdata edge"))

	return nil
}

func DocumentDataHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup documentdatahistory edge"))

	return nil
}

func EmailVerificationTokenEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup emailverificationtoken edge"))

	return nil
}

func EntitlementEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlement edge"))

	return nil
}

func EntitlementHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlementhistory edge"))

	return nil
}

func EntitlementPlanEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlementplan edge"))

	return nil
}

func EntitlementPlanFeatureEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlementplanfeature edge"))

	return nil
}

func EntitlementPlanFeatureHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlementplanfeaturehistory edge"))

	return nil
}

func EntitlementPlanHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitlementplanhistory edge"))

	return nil
}

func EntityEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entity edge"))

	return nil
}

func EntityHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entityhistory edge"))

	return nil
}

func EntityTypeEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitytype edge"))

	return nil
}

func EntityTypeHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup entitytypehistory edge"))

	return nil
}

func EventEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup event edge"))

	return nil
}

func EventHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup eventhistory edge"))

	return nil
}

func FeatureEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup feature edge"))

	return nil
}

func FeatureHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup featurehistory edge"))

	return nil
}

func FileEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup file edge"))

	return nil
}

func FileHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup filehistory edge"))

	return nil
}

func GroupEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup group edge"))

	if exists, err := FromContext(ctx).GroupSetting.Query().Where((groupsetting.HasGroupWith(group.ID(id)))).Exist(ctx); err == nil && exists {
		if groupsettingCount, err := FromContext(ctx).GroupSetting.Delete().Where(groupsetting.HasGroupWith(group.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", groupsettingCount).Msg("deleting groupsetting")
			return err
		}
	}

	if exists, err := FromContext(ctx).GroupMembership.Query().Where((groupmembership.HasGroupWith(group.ID(id)))).Exist(ctx); err == nil && exists {
		if groupmembershipCount, err := FromContext(ctx).GroupMembership.Delete().Where(groupmembership.HasGroupWith(group.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", groupmembershipCount).Msg("deleting groupmembership")
			return err
		}
	}

	return nil
}

func GroupHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup grouphistory edge"))

	return nil
}

func GroupMembershipEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup groupmembership edge"))

	return nil
}

func GroupMembershipHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup groupmembershiphistory edge"))

	return nil
}

func GroupSettingEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup groupsetting edge"))

	return nil
}

func GroupSettingHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup groupsettinghistory edge"))

	return nil
}

func HushEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup hush edge"))

	return nil
}

func HushHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup hushhistory edge"))

	return nil
}

func IntegrationEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup integration edge"))

	return nil
}

func IntegrationHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup integrationhistory edge"))

	return nil
}

func InternalPolicyEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup internalpolicy edge"))

	return nil
}

func InternalPolicyHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup internalpolicyhistory edge"))

	return nil
}

func InviteEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup invite edge"))

	return nil
}

func NarrativeEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup narrative edge"))

	return nil
}

func NarrativeHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup narrativehistory edge"))

	return nil
}

func NoteEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup note edge"))

	return nil
}

func NoteHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup notehistory edge"))

	return nil
}

func OauthProviderEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup oauthprovider edge"))

	return nil
}

func OauthProviderHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup oauthproviderhistory edge"))

	return nil
}

func OhAuthTooTokenEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup ohauthtootoken edge"))

	return nil
}

func OrgMembershipEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup orgmembership edge"))

	return nil
}

func OrgMembershipHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup orgmembershiphistory edge"))

	return nil
}

func OrganizationEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup organization edge"))

	if exists, err := FromContext(ctx).Organization.Query().Where(organization.HasParentWith(organization.ID(id))).Exist(ctx); err == nil && exists {
		if organizationCount, err := FromContext(ctx).Organization.Delete().Where(organization.HasParentWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", organizationCount).Msg("deleting child organization")
			return err
		}
	}

	if exists, err := FromContext(ctx).Group.Query().Where((group.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if groupCount, err := FromContext(ctx).Group.Delete().Where(group.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", groupCount).Msg("deleting group")
			return err
		}
	}

	if exists, err := FromContext(ctx).Template.Query().Where((template.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if templateCount, err := FromContext(ctx).Template.Delete().Where(template.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", templateCount).Msg("deleting template")
			return err
		}
	}

	if exists, err := FromContext(ctx).Integration.Query().Where((integration.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if integrationCount, err := FromContext(ctx).Integration.Delete().Where(integration.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", integrationCount).Msg("deleting integration")
			return err
		}
	}

	if exists, err := FromContext(ctx).OrganizationSetting.Query().Where((organizationsetting.HasOrganizationWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if organizationsettingCount, err := FromContext(ctx).OrganizationSetting.Delete().Where(organizationsetting.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", organizationsettingCount).Msg("deleting organizationsetting")
			return err
		}
	}

	if exists, err := FromContext(ctx).DocumentData.Query().Where((documentdata.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if documentdataCount, err := FromContext(ctx).DocumentData.Delete().Where(documentdata.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", documentdataCount).Msg("deleting documentdata")
			return err
		}
	}

	if exists, err := FromContext(ctx).Entitlement.Query().Where((entitlement.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if entitlementCount, err := FromContext(ctx).Entitlement.Delete().Where(entitlement.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", entitlementCount).Msg("deleting entitlement")
			return err
		}
	}

	if exists, err := FromContext(ctx).Invite.Query().Where((invite.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if inviteCount, err := FromContext(ctx).Invite.Delete().Where(invite.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", inviteCount).Msg("deleting invite")
			return err
		}
	}

	if exists, err := FromContext(ctx).Subscriber.Query().Where((subscriber.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if subscriberCount, err := FromContext(ctx).Subscriber.Delete().Where(subscriber.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", subscriberCount).Msg("deleting subscriber")
			return err
		}
	}

	if exists, err := FromContext(ctx).Webhook.Query().Where((webhook.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if webhookCount, err := FromContext(ctx).Webhook.Delete().Where(webhook.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", webhookCount).Msg("deleting webhook")
			return err
		}
	}

	if exists, err := FromContext(ctx).EntitlementPlan.Query().Where((entitlementplan.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if entitlementplanCount, err := FromContext(ctx).EntitlementPlan.Delete().Where(entitlementplan.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", entitlementplanCount).Msg("deleting entitlementplan")
			return err
		}
	}

	if exists, err := FromContext(ctx).EntitlementPlanFeature.Query().Where((entitlementplanfeature.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if entitlementplanfeatureCount, err := FromContext(ctx).EntitlementPlanFeature.Delete().Where(entitlementplanfeature.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", entitlementplanfeatureCount).Msg("deleting entitlementplanfeature")
			return err
		}
	}

	if exists, err := FromContext(ctx).Entity.Query().Where((entity.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if entityCount, err := FromContext(ctx).Entity.Delete().Where(entity.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", entityCount).Msg("deleting entity")
			return err
		}
	}

	if exists, err := FromContext(ctx).EntityType.Query().Where((entitytype.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if entitytypeCount, err := FromContext(ctx).EntityType.Delete().Where(entitytype.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", entitytypeCount).Msg("deleting entitytype")
			return err
		}
	}

	if exists, err := FromContext(ctx).Contact.Query().Where((contact.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if contactCount, err := FromContext(ctx).Contact.Delete().Where(contact.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", contactCount).Msg("deleting contact")
			return err
		}
	}

	if exists, err := FromContext(ctx).Note.Query().Where((note.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if noteCount, err := FromContext(ctx).Note.Delete().Where(note.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", noteCount).Msg("deleting note")
			return err
		}
	}

	if exists, err := FromContext(ctx).Task.Query().Where((task.HasOrganizationWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if taskCount, err := FromContext(ctx).Task.Delete().Where(task.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", taskCount).Msg("deleting task")
			return err
		}
	}

	if exists, err := FromContext(ctx).Program.Query().Where((program.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if programCount, err := FromContext(ctx).Program.Delete().Where(program.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", programCount).Msg("deleting program")
			return err
		}
	}

	if exists, err := FromContext(ctx).OrgMembership.Query().Where((orgmembership.HasOrganizationWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if orgmembershipCount, err := FromContext(ctx).OrgMembership.Delete().Where(orgmembership.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", orgmembershipCount).Msg("deleting orgmembership")
			return err
		}
	}

	return nil
}

func OrganizationHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup organizationhistory edge"))

	return nil
}

func OrganizationSettingEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup organizationsetting edge"))

	return nil
}

func OrganizationSettingHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup organizationsettinghistory edge"))

	return nil
}

func PasswordResetTokenEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup passwordresettoken edge"))

	return nil
}

func PersonalAccessTokenEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup personalaccesstoken edge"))

	return nil
}

func ProcedureEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup procedure edge"))

	return nil
}

func ProcedureHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup procedurehistory edge"))

	return nil
}

func ProgramEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup program edge"))

	if exists, err := FromContext(ctx).ProgramMembership.Query().Where((programmembership.HasProgramWith(program.ID(id)))).Exist(ctx); err == nil && exists {
		if programmembershipCount, err := FromContext(ctx).ProgramMembership.Delete().Where(programmembership.HasProgramWith(program.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", programmembershipCount).Msg("deleting programmembership")
			return err
		}
	}

	return nil
}

func ProgramHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup programhistory edge"))

	return nil
}

func ProgramMembershipEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup programmembership edge"))

	return nil
}

func ProgramMembershipHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup programmembershiphistory edge"))

	return nil
}

func RiskEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup risk edge"))

	return nil
}

func RiskHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup riskhistory edge"))

	return nil
}

func StandardEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup standard edge"))

	return nil
}

func StandardHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup standardhistory edge"))

	return nil
}

func SubcontrolEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup subcontrol edge"))

	return nil
}

func SubcontrolHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup subcontrolhistory edge"))

	return nil
}

func SubscriberEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup subscriber edge"))

	return nil
}

func TFASettingEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup tfasetting edge"))

	return nil
}

func TaskEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup task edge"))

	return nil
}

func TaskHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup taskhistory edge"))

	return nil
}

func TemplateEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup template edge"))

	if exists, err := FromContext(ctx).DocumentData.Query().Where((documentdata.HasTemplateWith(template.ID(id)))).Exist(ctx); err == nil && exists {
		if documentdataCount, err := FromContext(ctx).DocumentData.Delete().Where(documentdata.HasTemplateWith(template.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", documentdataCount).Msg("deleting documentdata")
			return err
		}
	}

	return nil
}

func TemplateHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup templatehistory edge"))

	return nil
}

func UserEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup user edge"))

	if exists, err := FromContext(ctx).PersonalAccessToken.Query().Where((personalaccesstoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if personalaccesstokenCount, err := FromContext(ctx).PersonalAccessToken.Delete().Where(personalaccesstoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", personalaccesstokenCount).Msg("deleting personalaccesstoken")
			return err
		}
	}

	if exists, err := FromContext(ctx).TFASetting.Query().Where((tfasetting.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if tfasettingCount, err := FromContext(ctx).TFASetting.Delete().Where(tfasetting.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", tfasettingCount).Msg("deleting tfasetting")
			return err
		}
	}

	if exists, err := FromContext(ctx).UserSetting.Query().Where((usersetting.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if usersettingCount, err := FromContext(ctx).UserSetting.Delete().Where(usersetting.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", usersettingCount).Msg("deleting usersetting")
			return err
		}
	}

	if exists, err := FromContext(ctx).EmailVerificationToken.Query().Where((emailverificationtoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if emailverificationtokenCount, err := FromContext(ctx).EmailVerificationToken.Delete().Where(emailverificationtoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", emailverificationtokenCount).Msg("deleting emailverificationtoken")
			return err
		}
	}

	if exists, err := FromContext(ctx).PasswordResetToken.Query().Where((passwordresettoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if passwordresettokenCount, err := FromContext(ctx).PasswordResetToken.Delete().Where(passwordresettoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", passwordresettokenCount).Msg("deleting passwordresettoken")
			return err
		}
	}

	if exists, err := FromContext(ctx).Webauthn.Query().Where((webauthn.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if webauthnCount, err := FromContext(ctx).Webauthn.Delete().Where(webauthn.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", webauthnCount).Msg("deleting webauthn")
			return err
		}
	}

	if exists, err := FromContext(ctx).OrgMembership.Query().Where((orgmembership.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if orgmembershipCount, err := FromContext(ctx).OrgMembership.Delete().Where(orgmembership.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", orgmembershipCount).Msg("deleting orgmembership")
			return err
		}
	}

	if exists, err := FromContext(ctx).GroupMembership.Query().Where((groupmembership.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if groupmembershipCount, err := FromContext(ctx).GroupMembership.Delete().Where(groupmembership.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			log.Debug().Err(err).Int("count", groupmembershipCount).Msg("deleting groupmembership")
			return err
		}
	}

	return nil
}

func UserHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup userhistory edge"))

	return nil
}

func UserSettingEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup usersetting edge"))

	return nil
}

func UserSettingHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup usersettinghistory edge"))

	return nil
}

func WebauthnEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup webauthn edge"))

	return nil
}

func WebhookEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup webhook edge"))

	return nil
}

func WebhookHistoryEdgeCleanup(ctx context.Context, id string) error {
	// If a user has access to delete the object, they have access to delete all edges
	ctx = privacy.DecisionContext(ctx, privacy.Allowf("cleanup webhookhistory edge"))

	return nil
}
