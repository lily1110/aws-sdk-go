package iam_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/iam"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleIAM_AddClientIDToOpenIDConnectProvider() {
	svc := iam.New(nil)

	params := &iam.AddClientIDToOpenIDConnectProviderInput{
		ClientID:                 aws.String("clientIDType"), // Required
		OpenIDConnectProviderARN: aws.String("arnType"),      // Required
	}
	resp, err := svc.AddClientIDToOpenIDConnectProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_AddRoleToInstanceProfile() {
	svc := iam.New(nil)

	params := &iam.AddRoleToInstanceProfileInput{
		InstanceProfileName: aws.String("instanceProfileNameType"), // Required
		RoleName:            aws.String("roleNameType"),            // Required
	}
	resp, err := svc.AddRoleToInstanceProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_AddUserToGroup() {
	svc := iam.New(nil)

	params := &iam.AddUserToGroupInput{
		GroupName: aws.String("groupNameType"),        // Required
		UserName:  aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.AddUserToGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_AttachGroupPolicy() {
	svc := iam.New(nil)

	params := &iam.AttachGroupPolicyInput{
		GroupName: aws.String("groupNameType"), // Required
		PolicyARN: aws.String("arnType"),       // Required
	}
	resp, err := svc.AttachGroupPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_AttachRolePolicy() {
	svc := iam.New(nil)

	params := &iam.AttachRolePolicyInput{
		PolicyARN: aws.String("arnType"),      // Required
		RoleName:  aws.String("roleNameType"), // Required
	}
	resp, err := svc.AttachRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_AttachUserPolicy() {
	svc := iam.New(nil)

	params := &iam.AttachUserPolicyInput{
		PolicyARN: aws.String("arnType"),      // Required
		UserName:  aws.String("userNameType"), // Required
	}
	resp, err := svc.AttachUserPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ChangePassword() {
	svc := iam.New(nil)

	params := &iam.ChangePasswordInput{
		NewPassword: aws.String("passwordType"), // Required
		OldPassword: aws.String("passwordType"), // Required
	}
	resp, err := svc.ChangePassword(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateAccessKey() {
	svc := iam.New(nil)

	params := &iam.CreateAccessKeyInput{
		UserName: aws.String("existingUserNameType"),
	}
	resp, err := svc.CreateAccessKey(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateAccountAlias() {
	svc := iam.New(nil)

	params := &iam.CreateAccountAliasInput{
		AccountAlias: aws.String("accountAliasType"), // Required
	}
	resp, err := svc.CreateAccountAlias(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateGroup() {
	svc := iam.New(nil)

	params := &iam.CreateGroupInput{
		GroupName: aws.String("groupNameType"), // Required
		Path:      aws.String("pathType"),
	}
	resp, err := svc.CreateGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateInstanceProfile() {
	svc := iam.New(nil)

	params := &iam.CreateInstanceProfileInput{
		InstanceProfileName: aws.String("instanceProfileNameType"), // Required
		Path:                aws.String("pathType"),
	}
	resp, err := svc.CreateInstanceProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateLoginProfile() {
	svc := iam.New(nil)

	params := &iam.CreateLoginProfileInput{
		Password:              aws.String("passwordType"), // Required
		UserName:              aws.String("userNameType"), // Required
		PasswordResetRequired: aws.Boolean(true),
	}
	resp, err := svc.CreateLoginProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateOpenIDConnectProvider() {
	svc := iam.New(nil)

	params := &iam.CreateOpenIDConnectProviderInput{
		ThumbprintList: []*string{ // Required
			aws.String("thumbprintType"), // Required
			// More values...
		},
		URL: aws.String("OpenIDConnectProviderUrlType"), // Required
		ClientIDList: []*string{
			aws.String("clientIDType"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateOpenIDConnectProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreatePolicy() {
	svc := iam.New(nil)

	params := &iam.CreatePolicyInput{
		PolicyDocument: aws.String("policyDocumentType"), // Required
		PolicyName:     aws.String("policyNameType"),     // Required
		Description:    aws.String("policyDescriptionType"),
		Path:           aws.String("policyPathType"),
	}
	resp, err := svc.CreatePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreatePolicyVersion() {
	svc := iam.New(nil)

	params := &iam.CreatePolicyVersionInput{
		PolicyARN:      aws.String("arnType"),            // Required
		PolicyDocument: aws.String("policyDocumentType"), // Required
		SetAsDefault:   aws.Boolean(true),
	}
	resp, err := svc.CreatePolicyVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateRole() {
	svc := iam.New(nil)

	params := &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String("policyDocumentType"), // Required
		RoleName:                 aws.String("roleNameType"),       // Required
		Path:                     aws.String("pathType"),
	}
	resp, err := svc.CreateRole(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateSAMLProvider() {
	svc := iam.New(nil)

	params := &iam.CreateSAMLProviderInput{
		Name:                 aws.String("SAMLProviderNameType"),     // Required
		SAMLMetadataDocument: aws.String("SAMLMetadataDocumentType"), // Required
	}
	resp, err := svc.CreateSAMLProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateUser() {
	svc := iam.New(nil)

	params := &iam.CreateUserInput{
		UserName: aws.String("userNameType"), // Required
		Path:     aws.String("pathType"),
	}
	resp, err := svc.CreateUser(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_CreateVirtualMFADevice() {
	svc := iam.New(nil)

	params := &iam.CreateVirtualMFADeviceInput{
		VirtualMFADeviceName: aws.String("virtualMFADeviceName"), // Required
		Path:                 aws.String("pathType"),
	}
	resp, err := svc.CreateVirtualMFADevice(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeactivateMFADevice() {
	svc := iam.New(nil)

	params := &iam.DeactivateMFADeviceInput{
		SerialNumber: aws.String("serialNumberType"),     // Required
		UserName:     aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.DeactivateMFADevice(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteAccessKey() {
	svc := iam.New(nil)

	params := &iam.DeleteAccessKeyInput{
		AccessKeyID: aws.String("accessKeyIdType"), // Required
		UserName:    aws.String("existingUserNameType"),
	}
	resp, err := svc.DeleteAccessKey(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteAccountAlias() {
	svc := iam.New(nil)

	params := &iam.DeleteAccountAliasInput{
		AccountAlias: aws.String("accountAliasType"), // Required
	}
	resp, err := svc.DeleteAccountAlias(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteAccountPasswordPolicy() {
	svc := iam.New(nil)

	var params *iam.DeleteAccountPasswordPolicyInput
	resp, err := svc.DeleteAccountPasswordPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteGroup() {
	svc := iam.New(nil)

	params := &iam.DeleteGroupInput{
		GroupName: aws.String("groupNameType"), // Required
	}
	resp, err := svc.DeleteGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteGroupPolicy() {
	svc := iam.New(nil)

	params := &iam.DeleteGroupPolicyInput{
		GroupName:  aws.String("groupNameType"),  // Required
		PolicyName: aws.String("policyNameType"), // Required
	}
	resp, err := svc.DeleteGroupPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteInstanceProfile() {
	svc := iam.New(nil)

	params := &iam.DeleteInstanceProfileInput{
		InstanceProfileName: aws.String("instanceProfileNameType"), // Required
	}
	resp, err := svc.DeleteInstanceProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteLoginProfile() {
	svc := iam.New(nil)

	params := &iam.DeleteLoginProfileInput{
		UserName: aws.String("userNameType"), // Required
	}
	resp, err := svc.DeleteLoginProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteOpenIDConnectProvider() {
	svc := iam.New(nil)

	params := &iam.DeleteOpenIDConnectProviderInput{
		OpenIDConnectProviderARN: aws.String("arnType"), // Required
	}
	resp, err := svc.DeleteOpenIDConnectProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeletePolicy() {
	svc := iam.New(nil)

	params := &iam.DeletePolicyInput{
		PolicyARN: aws.String("arnType"), // Required
	}
	resp, err := svc.DeletePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeletePolicyVersion() {
	svc := iam.New(nil)

	params := &iam.DeletePolicyVersionInput{
		PolicyARN: aws.String("arnType"),             // Required
		VersionID: aws.String("policyVersionIdType"), // Required
	}
	resp, err := svc.DeletePolicyVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteRole() {
	svc := iam.New(nil)

	params := &iam.DeleteRoleInput{
		RoleName: aws.String("roleNameType"), // Required
	}
	resp, err := svc.DeleteRole(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteRolePolicy() {
	svc := iam.New(nil)

	params := &iam.DeleteRolePolicyInput{
		PolicyName: aws.String("policyNameType"), // Required
		RoleName:   aws.String("roleNameType"),   // Required
	}
	resp, err := svc.DeleteRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteSAMLProvider() {
	svc := iam.New(nil)

	params := &iam.DeleteSAMLProviderInput{
		SAMLProviderARN: aws.String("arnType"), // Required
	}
	resp, err := svc.DeleteSAMLProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteServerCertificate() {
	svc := iam.New(nil)

	params := &iam.DeleteServerCertificateInput{
		ServerCertificateName: aws.String("serverCertificateNameType"), // Required
	}
	resp, err := svc.DeleteServerCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteSigningCertificate() {
	svc := iam.New(nil)

	params := &iam.DeleteSigningCertificateInput{
		CertificateID: aws.String("certificateIdType"), // Required
		UserName:      aws.String("existingUserNameType"),
	}
	resp, err := svc.DeleteSigningCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteUser() {
	svc := iam.New(nil)

	params := &iam.DeleteUserInput{
		UserName: aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.DeleteUser(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteUserPolicy() {
	svc := iam.New(nil)

	params := &iam.DeleteUserPolicyInput{
		PolicyName: aws.String("policyNameType"),       // Required
		UserName:   aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.DeleteUserPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DeleteVirtualMFADevice() {
	svc := iam.New(nil)

	params := &iam.DeleteVirtualMFADeviceInput{
		SerialNumber: aws.String("serialNumberType"), // Required
	}
	resp, err := svc.DeleteVirtualMFADevice(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DetachGroupPolicy() {
	svc := iam.New(nil)

	params := &iam.DetachGroupPolicyInput{
		GroupName: aws.String("groupNameType"), // Required
		PolicyARN: aws.String("arnType"),       // Required
	}
	resp, err := svc.DetachGroupPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DetachRolePolicy() {
	svc := iam.New(nil)

	params := &iam.DetachRolePolicyInput{
		PolicyARN: aws.String("arnType"),      // Required
		RoleName:  aws.String("roleNameType"), // Required
	}
	resp, err := svc.DetachRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_DetachUserPolicy() {
	svc := iam.New(nil)

	params := &iam.DetachUserPolicyInput{
		PolicyARN: aws.String("arnType"),      // Required
		UserName:  aws.String("userNameType"), // Required
	}
	resp, err := svc.DetachUserPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_EnableMFADevice() {
	svc := iam.New(nil)

	params := &iam.EnableMFADeviceInput{
		AuthenticationCode1: aws.String("authenticationCodeType"), // Required
		AuthenticationCode2: aws.String("authenticationCodeType"), // Required
		SerialNumber:        aws.String("serialNumberType"),       // Required
		UserName:            aws.String("existingUserNameType"),   // Required
	}
	resp, err := svc.EnableMFADevice(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GenerateCredentialReport() {
	svc := iam.New(nil)

	var params *iam.GenerateCredentialReportInput
	resp, err := svc.GenerateCredentialReport(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetAccessKeyLastUsed() {
	svc := iam.New(nil)

	params := &iam.GetAccessKeyLastUsedInput{
		AccessKeyID: aws.String("accessKeyIdType"), // Required
	}
	resp, err := svc.GetAccessKeyLastUsed(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetAccountAuthorizationDetails() {
	svc := iam.New(nil)

	params := &iam.GetAccountAuthorizationDetailsInput{
		Filter: []*string{
			aws.String("EntityType"), // Required
			// More values...
		},
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.GetAccountAuthorizationDetails(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetAccountPasswordPolicy() {
	svc := iam.New(nil)

	var params *iam.GetAccountPasswordPolicyInput
	resp, err := svc.GetAccountPasswordPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetAccountSummary() {
	svc := iam.New(nil)

	var params *iam.GetAccountSummaryInput
	resp, err := svc.GetAccountSummary(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetCredentialReport() {
	svc := iam.New(nil)

	var params *iam.GetCredentialReportInput
	resp, err := svc.GetCredentialReport(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetGroup() {
	svc := iam.New(nil)

	params := &iam.GetGroupInput{
		GroupName: aws.String("groupNameType"), // Required
		Marker:    aws.String("markerType"),
		MaxItems:  aws.Long(1),
	}
	resp, err := svc.GetGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetGroupPolicy() {
	svc := iam.New(nil)

	params := &iam.GetGroupPolicyInput{
		GroupName:  aws.String("groupNameType"),  // Required
		PolicyName: aws.String("policyNameType"), // Required
	}
	resp, err := svc.GetGroupPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetInstanceProfile() {
	svc := iam.New(nil)

	params := &iam.GetInstanceProfileInput{
		InstanceProfileName: aws.String("instanceProfileNameType"), // Required
	}
	resp, err := svc.GetInstanceProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetLoginProfile() {
	svc := iam.New(nil)

	params := &iam.GetLoginProfileInput{
		UserName: aws.String("userNameType"), // Required
	}
	resp, err := svc.GetLoginProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetOpenIDConnectProvider() {
	svc := iam.New(nil)

	params := &iam.GetOpenIDConnectProviderInput{
		OpenIDConnectProviderARN: aws.String("arnType"), // Required
	}
	resp, err := svc.GetOpenIDConnectProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetPolicy() {
	svc := iam.New(nil)

	params := &iam.GetPolicyInput{
		PolicyARN: aws.String("arnType"), // Required
	}
	resp, err := svc.GetPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetPolicyVersion() {
	svc := iam.New(nil)

	params := &iam.GetPolicyVersionInput{
		PolicyARN: aws.String("arnType"),             // Required
		VersionID: aws.String("policyVersionIdType"), // Required
	}
	resp, err := svc.GetPolicyVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetRole() {
	svc := iam.New(nil)

	params := &iam.GetRoleInput{
		RoleName: aws.String("roleNameType"), // Required
	}
	resp, err := svc.GetRole(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetRolePolicy() {
	svc := iam.New(nil)

	params := &iam.GetRolePolicyInput{
		PolicyName: aws.String("policyNameType"), // Required
		RoleName:   aws.String("roleNameType"),   // Required
	}
	resp, err := svc.GetRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetSAMLProvider() {
	svc := iam.New(nil)

	params := &iam.GetSAMLProviderInput{
		SAMLProviderARN: aws.String("arnType"), // Required
	}
	resp, err := svc.GetSAMLProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetServerCertificate() {
	svc := iam.New(nil)

	params := &iam.GetServerCertificateInput{
		ServerCertificateName: aws.String("serverCertificateNameType"), // Required
	}
	resp, err := svc.GetServerCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetUser() {
	svc := iam.New(nil)

	params := &iam.GetUserInput{
		UserName: aws.String("existingUserNameType"),
	}
	resp, err := svc.GetUser(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_GetUserPolicy() {
	svc := iam.New(nil)

	params := &iam.GetUserPolicyInput{
		PolicyName: aws.String("policyNameType"),       // Required
		UserName:   aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.GetUserPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListAccessKeys() {
	svc := iam.New(nil)

	params := &iam.ListAccessKeysInput{
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
		UserName: aws.String("existingUserNameType"),
	}
	resp, err := svc.ListAccessKeys(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListAccountAliases() {
	svc := iam.New(nil)

	params := &iam.ListAccountAliasesInput{
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListAccountAliases(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListAttachedGroupPolicies() {
	svc := iam.New(nil)

	params := &iam.ListAttachedGroupPoliciesInput{
		GroupName:  aws.String("groupNameType"), // Required
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("policyPathType"),
	}
	resp, err := svc.ListAttachedGroupPolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListAttachedRolePolicies() {
	svc := iam.New(nil)

	params := &iam.ListAttachedRolePoliciesInput{
		RoleName:   aws.String("roleNameType"), // Required
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("policyPathType"),
	}
	resp, err := svc.ListAttachedRolePolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListAttachedUserPolicies() {
	svc := iam.New(nil)

	params := &iam.ListAttachedUserPoliciesInput{
		UserName:   aws.String("userNameType"), // Required
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("policyPathType"),
	}
	resp, err := svc.ListAttachedUserPolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListEntitiesForPolicy() {
	svc := iam.New(nil)

	params := &iam.ListEntitiesForPolicyInput{
		PolicyARN:    aws.String("arnType"), // Required
		EntityFilter: aws.String("EntityType"),
		Marker:       aws.String("markerType"),
		MaxItems:     aws.Long(1),
		PathPrefix:   aws.String("pathType"),
	}
	resp, err := svc.ListEntitiesForPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListGroupPolicies() {
	svc := iam.New(nil)

	params := &iam.ListGroupPoliciesInput{
		GroupName: aws.String("groupNameType"), // Required
		Marker:    aws.String("markerType"),
		MaxItems:  aws.Long(1),
	}
	resp, err := svc.ListGroupPolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListGroups() {
	svc := iam.New(nil)

	params := &iam.ListGroupsInput{
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("pathPrefixType"),
	}
	resp, err := svc.ListGroups(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListGroupsForUser() {
	svc := iam.New(nil)

	params := &iam.ListGroupsForUserInput{
		UserName: aws.String("existingUserNameType"), // Required
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListGroupsForUser(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListInstanceProfiles() {
	svc := iam.New(nil)

	params := &iam.ListInstanceProfilesInput{
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("pathPrefixType"),
	}
	resp, err := svc.ListInstanceProfiles(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListInstanceProfilesForRole() {
	svc := iam.New(nil)

	params := &iam.ListInstanceProfilesForRoleInput{
		RoleName: aws.String("roleNameType"), // Required
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListInstanceProfilesForRole(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListMFADevices() {
	svc := iam.New(nil)

	params := &iam.ListMFADevicesInput{
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
		UserName: aws.String("existingUserNameType"),
	}
	resp, err := svc.ListMFADevices(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListOpenIDConnectProviders() {
	svc := iam.New(nil)

	var params *iam.ListOpenIDConnectProvidersInput
	resp, err := svc.ListOpenIDConnectProviders(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListPolicies() {
	svc := iam.New(nil)

	params := &iam.ListPoliciesInput{
		Marker:       aws.String("markerType"),
		MaxItems:     aws.Long(1),
		OnlyAttached: aws.Boolean(true),
		PathPrefix:   aws.String("policyPathType"),
		Scope:        aws.String("policyScopeType"),
	}
	resp, err := svc.ListPolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListPolicyVersions() {
	svc := iam.New(nil)

	params := &iam.ListPolicyVersionsInput{
		PolicyARN: aws.String("arnType"), // Required
		Marker:    aws.String("markerType"),
		MaxItems:  aws.Long(1),
	}
	resp, err := svc.ListPolicyVersions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListRolePolicies() {
	svc := iam.New(nil)

	params := &iam.ListRolePoliciesInput{
		RoleName: aws.String("roleNameType"), // Required
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListRolePolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListRoles() {
	svc := iam.New(nil)

	params := &iam.ListRolesInput{
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("pathPrefixType"),
	}
	resp, err := svc.ListRoles(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListSAMLProviders() {
	svc := iam.New(nil)

	var params *iam.ListSAMLProvidersInput
	resp, err := svc.ListSAMLProviders(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListServerCertificates() {
	svc := iam.New(nil)

	params := &iam.ListServerCertificatesInput{
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("pathPrefixType"),
	}
	resp, err := svc.ListServerCertificates(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListSigningCertificates() {
	svc := iam.New(nil)

	params := &iam.ListSigningCertificatesInput{
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
		UserName: aws.String("existingUserNameType"),
	}
	resp, err := svc.ListSigningCertificates(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListUserPolicies() {
	svc := iam.New(nil)

	params := &iam.ListUserPoliciesInput{
		UserName: aws.String("existingUserNameType"), // Required
		Marker:   aws.String("markerType"),
		MaxItems: aws.Long(1),
	}
	resp, err := svc.ListUserPolicies(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListUsers() {
	svc := iam.New(nil)

	params := &iam.ListUsersInput{
		Marker:     aws.String("markerType"),
		MaxItems:   aws.Long(1),
		PathPrefix: aws.String("pathPrefixType"),
	}
	resp, err := svc.ListUsers(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ListVirtualMFADevices() {
	svc := iam.New(nil)

	params := &iam.ListVirtualMFADevicesInput{
		AssignmentStatus: aws.String("assignmentStatusType"),
		Marker:           aws.String("markerType"),
		MaxItems:         aws.Long(1),
	}
	resp, err := svc.ListVirtualMFADevices(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_PutGroupPolicy() {
	svc := iam.New(nil)

	params := &iam.PutGroupPolicyInput{
		GroupName:      aws.String("groupNameType"),      // Required
		PolicyDocument: aws.String("policyDocumentType"), // Required
		PolicyName:     aws.String("policyNameType"),     // Required
	}
	resp, err := svc.PutGroupPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_PutRolePolicy() {
	svc := iam.New(nil)

	params := &iam.PutRolePolicyInput{
		PolicyDocument: aws.String("policyDocumentType"), // Required
		PolicyName:     aws.String("policyNameType"),     // Required
		RoleName:       aws.String("roleNameType"),       // Required
	}
	resp, err := svc.PutRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_PutUserPolicy() {
	svc := iam.New(nil)

	params := &iam.PutUserPolicyInput{
		PolicyDocument: aws.String("policyDocumentType"),   // Required
		PolicyName:     aws.String("policyNameType"),       // Required
		UserName:       aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.PutUserPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_RemoveClientIDFromOpenIDConnectProvider() {
	svc := iam.New(nil)

	params := &iam.RemoveClientIDFromOpenIDConnectProviderInput{
		ClientID:                 aws.String("clientIDType"), // Required
		OpenIDConnectProviderARN: aws.String("arnType"),      // Required
	}
	resp, err := svc.RemoveClientIDFromOpenIDConnectProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_RemoveRoleFromInstanceProfile() {
	svc := iam.New(nil)

	params := &iam.RemoveRoleFromInstanceProfileInput{
		InstanceProfileName: aws.String("instanceProfileNameType"), // Required
		RoleName:            aws.String("roleNameType"),            // Required
	}
	resp, err := svc.RemoveRoleFromInstanceProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_RemoveUserFromGroup() {
	svc := iam.New(nil)

	params := &iam.RemoveUserFromGroupInput{
		GroupName: aws.String("groupNameType"),        // Required
		UserName:  aws.String("existingUserNameType"), // Required
	}
	resp, err := svc.RemoveUserFromGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_ResyncMFADevice() {
	svc := iam.New(nil)

	params := &iam.ResyncMFADeviceInput{
		AuthenticationCode1: aws.String("authenticationCodeType"), // Required
		AuthenticationCode2: aws.String("authenticationCodeType"), // Required
		SerialNumber:        aws.String("serialNumberType"),       // Required
		UserName:            aws.String("existingUserNameType"),   // Required
	}
	resp, err := svc.ResyncMFADevice(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_SetDefaultPolicyVersion() {
	svc := iam.New(nil)

	params := &iam.SetDefaultPolicyVersionInput{
		PolicyARN: aws.String("arnType"),             // Required
		VersionID: aws.String("policyVersionIdType"), // Required
	}
	resp, err := svc.SetDefaultPolicyVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateAccessKey() {
	svc := iam.New(nil)

	params := &iam.UpdateAccessKeyInput{
		AccessKeyID: aws.String("accessKeyIdType"), // Required
		Status:      aws.String("statusType"),      // Required
		UserName:    aws.String("existingUserNameType"),
	}
	resp, err := svc.UpdateAccessKey(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateAccountPasswordPolicy() {
	svc := iam.New(nil)

	params := &iam.UpdateAccountPasswordPolicyInput{
		AllowUsersToChangePassword: aws.Boolean(true),
		HardExpiry:                 aws.Boolean(true),
		MaxPasswordAge:             aws.Long(1),
		MinimumPasswordLength:      aws.Long(1),
		PasswordReusePrevention:    aws.Long(1),
		RequireLowercaseCharacters: aws.Boolean(true),
		RequireNumbers:             aws.Boolean(true),
		RequireSymbols:             aws.Boolean(true),
		RequireUppercaseCharacters: aws.Boolean(true),
	}
	resp, err := svc.UpdateAccountPasswordPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateAssumeRolePolicy() {
	svc := iam.New(nil)

	params := &iam.UpdateAssumeRolePolicyInput{
		PolicyDocument: aws.String("policyDocumentType"), // Required
		RoleName:       aws.String("roleNameType"),       // Required
	}
	resp, err := svc.UpdateAssumeRolePolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateGroup() {
	svc := iam.New(nil)

	params := &iam.UpdateGroupInput{
		GroupName:    aws.String("groupNameType"), // Required
		NewGroupName: aws.String("groupNameType"),
		NewPath:      aws.String("pathType"),
	}
	resp, err := svc.UpdateGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateLoginProfile() {
	svc := iam.New(nil)

	params := &iam.UpdateLoginProfileInput{
		UserName:              aws.String("userNameType"), // Required
		Password:              aws.String("passwordType"),
		PasswordResetRequired: aws.Boolean(true),
	}
	resp, err := svc.UpdateLoginProfile(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateOpenIDConnectProviderThumbprint() {
	svc := iam.New(nil)

	params := &iam.UpdateOpenIDConnectProviderThumbprintInput{
		OpenIDConnectProviderARN: aws.String("arnType"), // Required
		ThumbprintList: []*string{ // Required
			aws.String("thumbprintType"), // Required
			// More values...
		},
	}
	resp, err := svc.UpdateOpenIDConnectProviderThumbprint(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateSAMLProvider() {
	svc := iam.New(nil)

	params := &iam.UpdateSAMLProviderInput{
		SAMLMetadataDocument: aws.String("SAMLMetadataDocumentType"), // Required
		SAMLProviderARN:      aws.String("arnType"),                  // Required
	}
	resp, err := svc.UpdateSAMLProvider(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateServerCertificate() {
	svc := iam.New(nil)

	params := &iam.UpdateServerCertificateInput{
		ServerCertificateName:    aws.String("serverCertificateNameType"), // Required
		NewPath:                  aws.String("pathType"),
		NewServerCertificateName: aws.String("serverCertificateNameType"),
	}
	resp, err := svc.UpdateServerCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateSigningCertificate() {
	svc := iam.New(nil)

	params := &iam.UpdateSigningCertificateInput{
		CertificateID: aws.String("certificateIdType"), // Required
		Status:        aws.String("statusType"),        // Required
		UserName:      aws.String("existingUserNameType"),
	}
	resp, err := svc.UpdateSigningCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UpdateUser() {
	svc := iam.New(nil)

	params := &iam.UpdateUserInput{
		UserName:    aws.String("existingUserNameType"), // Required
		NewPath:     aws.String("pathType"),
		NewUserName: aws.String("userNameType"),
	}
	resp, err := svc.UpdateUser(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UploadServerCertificate() {
	svc := iam.New(nil)

	params := &iam.UploadServerCertificateInput{
		CertificateBody:       aws.String("certificateBodyType"),       // Required
		PrivateKey:            aws.String("privateKeyType"),            // Required
		ServerCertificateName: aws.String("serverCertificateNameType"), // Required
		CertificateChain:      aws.String("certificateChainType"),
		Path:                  aws.String("pathType"),
	}
	resp, err := svc.UploadServerCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleIAM_UploadSigningCertificate() {
	svc := iam.New(nil)

	params := &iam.UploadSigningCertificateInput{
		CertificateBody: aws.String("certificateBodyType"), // Required
		UserName:        aws.String("existingUserNameType"),
	}
	resp, err := svc.UploadSigningCertificate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}