import { Command, program } from 'commander';
import open from 'open';
import pc from 'picocolors';
import jwtDecode from 'jwt-decode';
import inquirer from 'inquirer';
import { BaseCommandOptions } from '../../../core/types/types.js';
import { DecodedAccessToken, performDeviceAuth, startPollingForAccessToken } from '../utils.js';
import UserConfig from '../user-config.js';

export default (opts: BaseCommandOptions) => {
  const loginCommand = new Command('login');
  loginCommand.description('Login a user into the Cosmo platform. Supports browser-based authentication.');

  loginCommand.action(async () => {
    const userConfig = new UserConfig();
    // checks if the config file exists and checks if the access token is not expired
    userConfig.validateToken();

    const resp = await performDeviceAuth();
    if (!resp.success) {
      program.error('Could not perform authentication. Please try again');
    }
    console.log('Code: %s\n', resp.response.deviceCode);
    console.log(
      'If your browser does not automatically open, use this URL and enter the Code to verify the login: %s\n',
      resp.response.verificationURI,
    );
    console.log('Opening browser for login...\n');
    await open(resp.response.verificationURI);

    const accessTokenResp = await startPollingForAccessToken({
      deviceCode: resp.response.deviceCode,
      interval: resp.response.interval,
    });

    if (!accessTokenResp.success) {
      program.error(accessTokenResp.errorMessage + ' Please try again.');
    }

    if (!accessTokenResp.response) {
      program.error('Could not perform authentication. Please try again');
    }

    let decoded: DecodedAccessToken;

    try {
      decoded = jwtDecode<DecodedAccessToken>(accessTokenResp.response.accessToken);
    } catch {
      program.error('Could not perform authentication. Please try again');
    }

    const organizations = decoded.groups.map((group) => group.split('/')[1]);

    const selectedOrganization = await inquirer.prompt({
      name: 'organizationSlug',
      type: 'list',
      message: 'Select Organization:',
      choices: organizations,
    });

    userConfig.loadToken({ ...accessTokenResp.response, organizationSlug: selectedOrganization.organizationSlug });

    console.log(pc.green('Logged in Successfully!'));
  });

  return loginCommand;
};
